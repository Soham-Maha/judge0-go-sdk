package judge0

import "net/http"

type Client struct {
	authProvider AuthProvider
	httpClient   *http.Client
	baseURL      string
}
