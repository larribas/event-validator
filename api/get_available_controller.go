package api

import (
    "net/http"
    "github.com/sp-lorenzo-arribas/event_validator/domain"
    "fmt"
)

type GetAvailableResponse struct {
    MinVersion int `json:"min_version"`
    MaxVersion int `json:"max_version"`
}

func GetAvailableController(w http.ResponseWriter, r *http.Request) {
    _type := r.URL.Query().Get(":type")

    nextVersion := domain.Current.GetRepository().GetNextVersion(_type)
    if nextVersion == 0 {
        writeErrorMessage(w, "TypeWithNoValidators", fmt.Sprintf("Type '%s' has no available validators", _type))
        return
    }

    writeJSONResponse(w, GetAvailableResponse{0, nextVersion - 1})
}