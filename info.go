package judge0

import (
	"fmt"
	"net/http"
	"net/url"
	"io"
	"encoding/json"

)

type Status struct {
	Id       int  `json:"id"`
	Description string  `json:"description"`
}
type Statuses []Status

func Judge0Request(url string, method string) ([] byte, error) {
	judgereq, err := http.NewRequest(method, url, nil)
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


func (client * Client) StatusesGet() (Statuses, error)   {
	judge0Url, err := url.Parse(client.authProvider.GetBaseURL() + "/statuses")
	if err != nil {
		return nil, err
	}
	body, err := Judge0Request(judge0Url.String(), "GET")
	if err != nil {
		return nil, fmt.Errorf("error making : %v", err)
	}
	var submissions Statuses
	err = json.Unmarshal(body, &submissions)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %v", err)

	}
	return submissions, nil
}
