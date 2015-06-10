package api

import (
	"fmt"
	"github.com/sp-lorenzo-arribas/event_validator/domain"
	"net/http"
	"testing"
)

func TestInspect(t *testing.T) {
	ts := newTestServer()
	defer ts.Close()

	// Load fixtures
	validator, _ := domain.NewValidator("some-type", []byte(`{"type": "string"}`))
	version := domain.Current.GetRepository().Create(validator)

	url := fmt.Sprintf("%s/validators/%s/versions/%d", ts.URL, "some-type", version)
	res, err := http.Get(url)

	assertRightEndpoint("GET", url, err, t)
	assertStatusCode(res, 200, t)

	var respParameters InspectResponse
	checkAndUnmarshalJSON(res, &respParameters, t)
	if respParameters.Rules != string(validator.Rules) {
		t.Errorf("Expected GET /validators/some-type/versions/0 to respond with the right validator rules")
	}
}

func TestInspectNonexistent(t *testing.T) {
	ts := newTestServer()
	defer ts.Close()

	// Load fixtures
	validator, _ := domain.NewValidator("existing-type", []byte(`{"type": "string"}`))
	domain.Current.GetRepository().Create(validator)

	testCases := []struct {
		Type      string
		Version   string
		ErrorType string
	}{
		{"any-type", "nonnumeric-version", "InvalidParameter"},
		{"nonexistent-type", "0", "ValidatorDoesNotExist"},
		{"existing-type", "1212", "ValidatorDoesNotExist"},
	}

	for _, testCase := range testCases {
		url := fmt.Sprintf("%s/validators/%s/versions/%s", ts.URL, testCase.Type, testCase.Version)
		res, err := http.Get(url)

		assertRightEndpoint("GET", url, err, t)
		assertStatusCode(res, 400, t)
		assertErrorType(res, testCase.ErrorType, t)
	}
}
