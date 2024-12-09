package judge0

import (
	"fmt"
	"net/http"
)

// AuthProvider interface is used to set the authentication headers of all the requests
type AuthProvider interface {
	SetAuthHeaders(req *http.Request) // Set the Authentication Headers
	GetBaseURL() string               // Get the Base URL of the request
}

// Implements the AuthProivder Interface for the RapidAPI Judge0 Provider
type rapidAPIProivder struct {
	rapidAPIKey string
}

func (rap *rapidAPIProivder) SetAuthHeaders(req *http.Request) {
	req.Header.Set("x-rapidapi-key", rap.rapidAPIKey)
	req.Header.Set("x-rapidapi-host", rapidAPIHost)
}

func (rap *rapidAPIProivder) GetBaseURL() string {
	return rapidAPIHost
}

// Implements the AuthProvider Interface for the Sulu Judge0 Provider
type suluAPIProvider struct {
	suluAPIKey string
}

func (sap *suluAPIProvider) SetAuthHeaders(req *http.Request) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", sap.suluAPIKey))
}

func (sap *suluAPIProvider) GetBaseURL() string {
	return suluAPIHost
}

func NewSuluProivder(apiKey string) AuthProvider {
	return &suluAPIProvider{
		suluAPIKey: apiKey,
	}
}

func NewRapidAPIProvider(apiKey string) AuthProvider {
	return &rapidAPIProivder{
		rapidAPIKey: apiKey,
	}
}
