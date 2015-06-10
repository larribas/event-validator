package api

import (
	"fmt"
	"github.com/sp-lorenzo-arribas/event_validator/domain"
	"net/http"
)

// GetAvailableResponse represents this controller's response format
type GetAvailableResponse struct {
	MinVersion int `json:"min_version"`
	MaxVersion int `json:"max_version"`
}

// GetAvailableController handles a request to know the available versions for a certain event type. It returns a
// range of (min, max) versions. All versions in between are assumed to exist. If the specified type has no versions,
// it responds with a domain error
func GetAvailableController(w http.ResponseWriter, r *http.Request) {
	_type := r.URL.Query().Get(":type")

	nextVersion := domain.Current.GetRepository().GetNextVersion(_type)
	if nextVersion == 0 {
		writeErrorMessage(w, "TypeWithNoValidators", fmt.Sprintf("Type '%s' has no available validators", _type))
		return
	}

	writeJSONResponse(w, GetAvailableResponse{0, nextVersion - 1})
}
