package api

import (
    "net/http"
    "github.com/sp-lorenzo-arribas/event_validator/domain"
    "strconv"
)

type InspectResponse struct {
    Rules []byte `json:"rules"`
}

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