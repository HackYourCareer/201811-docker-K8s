/*
* Don't use in production...
 */

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mazehyc/web/redisHelper"
	"mazehyc/web/worker"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func main() {

	//@@ Reading config from environment variables
	port := os.Getenv("PORT") // port string
	if port == "" {
		port = "8081"
	}

	workRootDir := os.Getenv("WORKER_DIR")
	if workRootDir == "" {
		workRootDir = "./worker"
	}
	fmt.Printf("WORKER_DIR: %s\n", workRootDir)

	mazegenCommand := os.Getenv("MAZEGEN_CMD") //Maze generation command
	if mazegenCommand == "" {
		panic("MAZEGEN_CMD env variable is not set")
	}
	fmt.Printf("MAZEGEN_CMD: %s\n", mazegenCommand)

	//Redis Integration (optional for now)
	var storeInRedis redisHelper.SetRedisValueFunc = nil

	redisAddr := os.Getenv("REDIS_ADDR") //Redis Address
	if redisAddr == "" {
		fmt.Println("WARNING: REDIS_ADDR environment variable is missing, skipping integration with Redis!")
		storeInRedis = func(key, value string) error {
			//Do nothing at all
			return nil
		}
	} else {
		fmt.Printf("REDIS_ADDR: %s\n", redisAddr)
		redisClient := redisHelper.NewClient(redisAddr)
		storeInRedis = func(key, value string) error {
			fmt.Printf("Storing in Redis: %s => %s\n", key, value)
			return redisHelper.SetValue(redisClient, key, value, 3600)
		}
	}

	if _, err := os.Stat(workRootDir); os.IsNotExist(err) {
		err := os.Mkdir(workRootDir, os.ModePerm)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Created directory: %s\n", workRootDir)
	}

	//@@ Handler function (separate goroutines)
	http.HandleFunc("/generate", handleFunc(workRootDir, mazegenCommand, storeInRedis))

	fmt.Printf("Running server on port: %s\nType Ctr-c to shutdown server.\n", port)
	err := http.ListenAndServe(":"+port, nil)
	fmt.Println("err:\n  " + err.Error())
}

func handleFunc(workRootDir, mazegenCommand string, storeInRedis redisHelper.SetRedisValueFunc) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		url := r.URL
		params := url.Query()

		//@@ Input params (query string)
		token := params.Get("token")
		text := params.Get("text")
		width := params.Get("width")
		height := params.Get("height")

		if token == "" {
			http.Error(w, "Missing token", http.StatusBadRequest)
			return
		}

		if (text == "") && (width == "" || height == "") {
			http.Error(w, "Missing maze generation params (?text=... or ?width=...&height=...", http.StatusBadRequest)
			return
		}

		//@@ Absolute dir
		targetDir := workRootDir + "/" + token

		var args []string
		if text != "" {
			args = []string{targetDir, text}
		} else {
			args = []string{targetDir, width, height}
		}

		//@@ Executing command
		fmt.Printf("Executing command: %s, %#v\n", mazegenCommand, args)
		cmd := exec.Command(mazegenCommand, args...)

		var stderr bytes.Buffer
		var stdout bytes.Buffer
		cmd.Stderr = &stderr
		cmd.Stdout = &stdout
		err := cmd.Run()

		if err != nil {
			errStr := string(stderr.Bytes())
			outStr := string(stdout.Bytes())
			fmt.Println("Comand error:")
			fmt.Println(strings.Trim(errStr, "\n"))
			fmt.Println("Command output:")
			fmt.Println(strings.Trim(outStr, "\n"))
			http.Error(w, errStr, http.StatusBadRequest)
			return
		}

		//@@ Creating response to put in Redis
		res, _ := json.Marshal(worker.WorkerResponse{
			Ready:        true,
			ImageFile:    token + "/maze.original.png",
			EnhancedFile: token + "/maze.enhanced.png",
		})

		storeInRedis(token, string(res))

		//Log it
		fmt.Println(stdout.String())
		w.WriteHeader(http.StatusNoContent)
	}
}
