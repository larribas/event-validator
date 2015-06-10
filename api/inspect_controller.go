package api

import (
	"github.com/sp-lorenzo-arribas/event_validator/domain"
	"net/http"
	"strconv"
)

// InspectResponse represents this controller's response format
type InspectResponse struct {
	Rules []byte `json:"rules"`
}

// InspectController handles a request to get the validation rules for a certain (type, version) combination.
// If there is no such validator, it responds with a domain error. Otherwise, it returns the validation rules
// for such combination, represented as a string
func InspectController(w http.ResponseWriter, r *http.Request) {
	_type := r.URL.Query().Get(":type")
	version, err := strconv.Atoi(r.URL.Query().Get(":version"))
	if err != nil {
		writeErrorMessage(w, "InvalidParameter", "Parameter 'version' should be an integer")
		return
	}

	validator, err := domain.Current.GetRepository().Inspect(_type, version)
	if err != nil {
		writeErrorMessage(w, "ValidatorDoesNotExist", err.Error())
		return
	}

	writeJSONResponse(w, InspectResponse{validator.Rules})
}
