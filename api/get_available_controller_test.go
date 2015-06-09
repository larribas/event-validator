package api

import (
    "testing"
    "fmt"
    "github.com/sp-lorenzo-arribas/event_validator/domain"
    "net/http"
)

func TestGetAvailable(t *testing.T) {
    ts := newTestServer()
    defer ts.Close()

    // Load fixtures
    validator, _ := domain.NewValidator("some-type", []byte(`{"type": "string"}`))
    min := domain.Current.GetRepository().Create(validator)
    max := domain.Current.GetRepository().Create(validator)

    url := fmt.Sprintf("%s/validators/%s/versions", ts.URL, "some-type")
    res, err := http.Get(url)

    assertRightEndpoint("GET", url, err, t)
    assertStatusCode(res, 200, t)

    var respParameters GetAvailableResponse
    checkAndUnmarshalJSON(res, &respParameters, t)
    if respParameters.MinVersion != min || respParameters.MaxVersion != max {
        t.Errorf("Expected GET /validators/some-type/versions to respond with (%d, %d) interval of available versions. Instead, it responded with (%d, %d)", min, max, respParameters.MinVersion, respParameters.MaxVersion)
    }
}

func TestGetAvailableEmptyType(t *testing.T) {
    ts := newTestServer()
    defer ts.Close()

    url := fmt.Sprintf("%s/validators/%s/versions", ts.URL, "empty-type")
    res, err := http.Get(url)

    assertRightEndpoint("GET", url, err, t)
    assertStatusCode(res, 400, t)
    assertErrorType(res, "TypeWithNoValidators", t)
}