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


func (client * Client) StatusesGet() (Statuses, error)   {
	judge0Url, err := url.Parse(client.authProvider.GetBaseURL() + "/statuses")
	if err != nil {
		return nil, fmt.Errorf("error parsing Judge0 URL: %v", err)
	}
	judgereq, err := http.NewRequest("GET", judge0Url.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("error parsing Judge0 url to GET request: %v", err)
	}
	judgereq.Header.Add("Content-Type", "application/json")
	// judgereq.Header.Add("Authorization", fmt.Sprintf("Bearer %v", bearer))

	resp, err := http.DefaultClient.Do(judgereq)
	if err != nil {
		return nil, fmt.Errorf("error sending request to Judge0: %v", err)
		
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)

	}
	
	var submissions Statuses
	err = json.Unmarshal(body, &submissions)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %v", err)

	}
	return submissions, nil
}
