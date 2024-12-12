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
type rapidAPIProvider struct {
	rapidAPIKey string
}

func (rap *rapidAPIProvider) SetAuthHeaders(req *http.Request) {
	req.Header.Set("x-rapidapi-key", rap.rapidAPIKey)
	req.Header.Set("x-rapidapi-host", rapidAPIHost)
}

func (rap *rapidAPIProvider) GetBaseURL() string {
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

type customAPIProvider struct {
	baseURL      string
	headerField  string
	customAPIKey string
}

func (cap *customAPIProvider) SetAuthHeaders(req *http.Request) {
	req.Header.Set(cap.headerField, cap.customAPIKey)
}

func (cap *customAPIProvider) GetBaseURL() string {
	return cap.baseURL
}

func NewCustomProvider(apiKey string, baseUrl string, headerField string) AuthProvider {
	if headerField == "" {
		headerField = "X-Auth-Token"
	}
	return &customAPIProvider{
		baseURL:      baseUrl,
		headerField:  headerField,
		customAPIKey: apiKey,
	}
}

func NewSuluProvider(apiKey string) AuthProvider {
	return &suluAPIProvider{
		suluAPIKey: apiKey,
	}
}

func NewRapidAPIProvider(apiKey string) AuthProvider {
	return &rapidAPIProvider{
		rapidAPIKey: apiKey,
	}
}
