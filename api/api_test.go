package api

import (
	"encoding/json"
	"github.com/sp-lorenzo-arribas/event_validator/domain"
	"github.com/sp-lorenzo-arribas/event_validator/formats"
	"github.com/sp-lorenzo-arribas/event_validator/repositories"
	"net/http"
	"net/http/httptest"
	"testing"
)

// newTestServer is a helper function to initialize the domain with a lighweight testable environment
func newTestServer() *httptest.Server {
	// Set up a clean domain
	domain.Current = &domain.Environment{
		NewRepository: repositories.NewInMemoryRepository,
		NewFormatChecker: func() domain.FormatChecker {
			return &formats.JSONSchemaFormatChecker{}
		},
	}

	// Create a new API with zeroed address and port (httptest will use their own)
	api := New("", "")
	return httptest.NewServer(api.router)
}

// Generic Assertions

func assertRightEndpoint(method string, url string, err error, t *testing.T) {
	if err != nil {
		t.Errorf("Expected %s %s response not to return an error. Instead, it returned %s", method, url, err.Error())
	}
}

func assertStatusCode(r *http.Response, statusCode int, t *testing.T) {
	if r.StatusCode != statusCode {
		t.Errorf("Expected %s %s to return a %d status code. Instead, it returned %d", r.Request.Method, r.Request.URL.String(), statusCode, r.StatusCode)
	}
}

func assertErrorType(r *http.Response, _type string, t *testing.T) {
	defer r.Body.Close()
	var resp ErrorResponse
	json.NewDecoder(r.Body).Decode(&resp)
	if resp.Type != _type {
		t.Errorf("Expected %s %s to respond with an error of type '%s'. Instead, it responded '%s'", r.Request.Method, r.Request.URL.String(), _type, resp.Type)
	}
}

func checkAndUnmarshalJSON(r *http.Response, j interface{}, t *testing.T) {
	err := json.NewDecoder(r.Body).Decode(j)
	r.Body.Close()
	if err != nil {
		t.Errorf("Expected %s %s to return a properly formatted JSON response. Instead, unmarshaling returned %s", r.Request.Method, r.Request.URL.String(), err.Error())
	}
}
