package worker

type WorkerResponse struct {
	Ready        bool   `json:"ready"`
	ImageFile    string `json:"imageFile,omitempty"`
	EnhancedFile string `json:"printFile,omitempty"`
}
