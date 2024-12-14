package judge0

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Status struct {
	Id       int  `json:"id"`
	Description string  `json:"description"`
}
type Statuses []Status



func (client * Client) GetStatus() (Statuses, error)   {
	judge0Url, err := url.Parse(client.authProvider.GetBaseURL() + "/statuses")
	if err != nil {
		return nil, err
	}
	body, err := Judge0Request(judge0Url.String(), http.MethodGet, nil)
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
