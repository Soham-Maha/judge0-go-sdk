package judge0

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)
func Judge0Request(url string, method string, payload []byte) ([] byte, error) {
	judgereq, err := http.NewRequest(method, url, bytes.NewReader(payload))
	if err != nil {
		return nil, fmt.Errorf("error creating the HTTP request: %v", err)
	}
	judgereq.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(judgereq)
	if err != nil {
		return nil, fmt.Errorf("error sending request to Judge0: %v", err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	
	}
	return body, nil
}
