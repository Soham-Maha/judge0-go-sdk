package judge0

import "net/http"

type Client struct {
	authProvider AuthProvider
	httpClient   *http.Client
	baseURL      string
}

type Worker struct {
	Queue     string `json:"queue"`
	Size      int    `json:"size"`
	Available int    `json:"available"`
	Idle      int    `json:"idle"`
	Working   int    `json:"working"`
	Paused    int    `json:"paused"`
	Failed    int    `json:"failed"`
}

type Workers []Worker

func (client *Client) GetWorkers() (Workers, error) {

}
