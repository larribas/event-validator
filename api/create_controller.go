package api

import (
	"encoding/json"
	"github.com/sp-lorenzo-arribas/event_validator/domain"
	"net/http"
)

// CreateRequest represents this controller's mandatory request format
type CreateRequest struct {
	Rules *string `json:"rules"`
}

// CreateResponse represents this controller's response format
type CreateResponse struct {
	Version int `json:"version"`
}

// CreateController handles a request to create a new validator. It validates that the correct parameters are supplied,
// that the validator has the right format, and returns the version that corresponds to the new validator
func CreateController(w http.ResponseWriter, r *http.Request) {
	_type := r.URL.Query().Get(":type")

	var reqParams CreateRequest
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&reqParams)
	if err != nil || reqParams.Rules == nil {
		writeErrorMessage(w, "InvalidParameters", "The request should be formatted as application/json containing a 'rules' parameter")
		return
	}

	validator, err := domain.NewValidator(_type, []byte(*(reqParams.Rules)))
	if err != nil {
		writeErrorMessage(w, "ValidatorCouldNotBeCreated", err.Error())
		return
	}

	domain.Current.GetRepository().Create(validator)
	w.WriteHeader(http.StatusCreated)
	writeJSONResponse(w, CreateResponse{validator.Version})
}
