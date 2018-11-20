package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mazehyc/web/redisHelper"
	"mazehyc/web/worker"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/twinj/uuid"
)

func firstValue(vals []string) string {
	if len(vals) == 0 {
		return ""
	}

	return vals[0]
}

func main() {

	port := os.Getenv("PORT") // port string

	if port == "" {
		port = "8080"
	}

	workRootDir := os.Getenv("WORKER_DIR")
	if workRootDir == "" {
		workRootDir = "./generated"
	}
	fmt.Printf("WORK_ROOT_DIR: %s\n", workRootDir)

	workerHost := os.Getenv("WORKER_HOST")
	if workerHost == "" {
		panic("WORKER_HOST env variable is not set")
	}
	fmt.Printf("WORKER_HOST: %s\n", workerHost)

	//Redis Integration (optional for now)
	var getValueFunc redisHelper.GetRedisValueFunc = nil
	var setValueFunc redisHelper.SetRedisValueFunc = nil

	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		fmt.Println("WARNING: REDIS_ADDR environment variable is missing, skipping integration with Redis!")
		getValueFunc = func(key string) (bool, string, error) {
			//Do nothing at all
			return false, "", nil
		}
		setValueFunc = func(key, value string) error {
			//Do nothing at all
			return nil
		}
	} else {
		fmt.Printf("REDIS_ADDR: %s\n", redisAddr)
		redisClient := redisHelper.NewClient(redisAddr)
		getValueFunc = func(key string) (bool, string, error) {
			return redisHelper.GetValue(redisClient, key)
		}
		setValueFunc = func(key, value string) error {
			return redisHelper.SetValue(redisClient, key, value, 3600)
		}
	}

	//Handlers
	http.HandleFunc("/maze", mazeHandler(workerHost, setValueFunc))
	http.HandleFunc("/token/", tokenHandler(getValueFunc))
	http.HandleFunc("/result/", resultHandler(workRootDir, getValueFunc))

	fmt.Printf("Running server on port: %s\nType Ctr-c to shutdown server.\n", port)
	err := http.ListenAndServe(":"+port, nil)
	fmt.Println("err:\n  " + err.Error())
}

func mazeHandler(workerHost string, setValueFunc redisHelper.SetRedisValueFunc) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		params := r.URL.Query()

		//@@ Extract params from URL (query string)
		text := params.Get("text")
		width := params.Get("width")
		height := params.Get("height")

		if (text == "") && (width == "" || height == "") {
			http.Error(w, "Missing maze generation params (?text=... or ?width=...&height=...", http.StatusBadRequest)
			return
		}

		//@@ Generate token for the operation
		token := uuid.NewV4().String()

		//@@ Store initial data in Redis
		initialWorkerRes, _ := json.Marshal(worker.WorkerResponse{
			Ready: false,
		})
		err := setValueFunc(token, string(initialWorkerRes))

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		//@@ Prepare args for worker
		var args string
		if text != "" {
			args = "token=" + token + "&text=" + url.QueryEscape(text)
		} else {
			args = "token=" + token + "&width=" + width + "&height=" + height
		}

		workerUrl := "http://" + workerHost + "/generate?" + args
		fmt.Printf("Dispatching call to: %s\n", workerUrl)

		//@@ Dispatch call to worker
		workerResp, err := http.Get(workerUrl)

		//@@ Handle error
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		//@@ Handle Bad Request
		if workerResp.StatusCode == http.StatusBadRequest {
			//Log it
			defer workerResp.Body.Close()
			body, err := ioutil.ReadAll(workerResp.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Unexpected response from worker (%s):\n", workerResp.Status)
			fmt.Fprint(w, string(body))
			return
		}

		//@@ Handle unexpected status
		if workerResp.StatusCode != http.StatusNoContent {

			//Log it
			defer workerResp.Body.Close()
			body, err := ioutil.ReadAll(workerResp.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			http.Error(w, "Unexpected response from worker: "+workerResp.Status, http.StatusInternalServerError)
			fmt.Fprint(w, string(body))
			return
		}

		//@@ Send response to user
		mazeResp := mazeResponse{
			TokenUrl: "http://" + r.Host + "/token/" + token,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(mazeResp)
	}
}

func tokenHandler(getValueFunc redisHelper.GetRedisValueFunc) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		//@@ Extract token from path
		parts := strings.Split(r.URL.Path, "/token/")
		token := parts[len(parts)-1]

		//@@ Lookup for data in Redis
		found, redisValue, err := getValueFunc(token)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if !found {
			msg := "Result not ready yet"
			fmt.Println(msg)
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, msg)
			return
		}

		fileToUrlFunc := func(file string) string {
			//TODO: Convert
			return "http://" + r.Host + "/result/" + file
		}

		//@@ Convert from WorkerResponse to controllerResponse
		res, err := toControllerResponse(redisValue, fileToUrlFunc)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, res)
		return
	}
}

func resultHandler(workerRootDir string, getValueFunc redisHelper.GetRedisValueFunc) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		parts := strings.Split(r.URL.Path, "/result/")
		file := parts[len(parts)-1]

		path := workerRootDir + "/" + file
		fmt.Printf("Serving file: %s\n", path)
		http.ServeFile(w, r, path)
	}
}

// Converts from WorkerResponse (files) to controllerResponse (urls)
func toControllerResponse(redisValue string, fileFunc fileToUrlFunc) (string, error) {

	workerResp := worker.WorkerResponse{}
	err := json.Unmarshal([]byte(redisValue), &workerResp)

	if err != nil {
		return "", err
	}

	var imageUrl string
	var enhancedUrl string

	if workerResp.Ready {
		imageUrl = fileFunc(workerResp.ImageFile)
		enhancedUrl = fileFunc(workerResp.EnhancedFile)
	}

	controllerResp := controllerResponse{
		Ready:       workerResp.Ready,
		ImageUrl:    imageUrl,
		EnhancedUrl: enhancedUrl,
	}

	res, err := json.Marshal(controllerResp)
	if err != nil {
		return "", err
	}

	return string(res), nil
}

type fileToUrlFunc func(string) string

type mazeResponse struct {
	TokenUrl string `json:"tokenUrl"`
}

type tokenResponse struct {
	TokenUrl string `json:"tokenUrl"`
}

type controllerResponse struct {
	Ready       bool   `json:"ready"`
	ImageUrl    string `json:"imageUrl,omitempty"`
	EnhancedUrl string `json:"printUrl,omitempty"`
}
