package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestCreate(t *testing.T) {
	ts := newTestServer()
	defer ts.Close()

	url := fmt.Sprintf("%s/validators/%s", ts.URL, "some-type")
	req := CreateRequest{Rules: new(string)}
	*(req.Rules) = `{"type": "string"}`
	j, _ := json.Marshal(&req)

	for i := range []int{1, 2} {
		res, err := http.Post(url, "application/json", bytes.NewReader(j))

		assertRightEndpoint("POST", url, err, t)
		assertStatusCode(res, 201, t)

		var respParameters CreateResponse
		checkAndUnmarshalJSON(res, &respParameters, t)
		if respParameters.Version != i {
			t.Errorf("Expected POST %s to respond with version %d. Instead it responded with version %d", res.Request.URL.Path, i, respParameters.Version)
		}
	}
}

func TestCreateWithoutRules(t *testing.T) {
	ts := newTestServer()
	defer ts.Close()

	url := fmt.Sprintf("%s/validators/%s", ts.URL, "some-type")
	j, _ := json.Marshal("invalid request format")

	res, err := http.Post(url, "application/json", bytes.NewReader(j))

	assertRightEndpoint("POST", url, err, t)
	assertStatusCode(res, 400, t)
	assertErrorType(res, "InvalidParameters", t)
}

func TestCreateWrongValidator(t *testing.T) {
	ts := newTestServer()
	defer ts.Close()

	url := fmt.Sprintf("%s/validators/%s", ts.URL, "some-type")
	req := CreateRequest{Rules: new(string)}
	*(req.Rules) = `["invalid", "jsonschema"]`
	j, err := json.Marshal(&req)

	res, err := http.Post(url, "application/json", bytes.NewReader(j))

	assertRightEndpoint("POST", url, err, t)
	assertStatusCode(res, 400, t)
	assertErrorType(res, "ValidatorCouldNotBeCreated", t)
}
