package api

import (
	"encoding/json"
	"fmt"
	"github.com/bmizerany/pat"
	"log"
	"net/http"
)

// HTTPApi holds the components and behavior that allows for the API to be exposed
type HTTPApi struct {
	address, port string
	router        *pat.PatternServeMux
}

// New receives a address and a port, and sets the API ready to be exposed
func New(address, port string) *HTTPApi {
	api := &HTTPApi{address: address, port: port}
	api.router = pat.New()

	// Endpoints
	api.router.Post("/validators/:type", http.HandlerFunc(CreateController))
	api.router.Get("/validators/:type/versions", http.HandlerFunc(GetAvailableController))
	api.router.Get("/validators/:type/versions/:version", http.HandlerFunc(InspectController))

	return api
}

// Expose starts the API by listening on the address and port specified, and directing such requests to the
// specified routes and controllers
func (api *HTTPApi) Expose() {
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", api.address, api.port), api.router)
	if err != nil {
		log.Fatalln("[HTTP API] Couldn't listen and serve due to", err.Error())
	}
}

// HELPER METHODS TO HANDLE RESPONSES

type ErrorResponse struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

func writeJSONResponse(w http.ResponseWriter, j interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(j)
}

func writeErrorMessage(w http.ResponseWriter, _type, msg string) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(ErrorResponse{_type, msg})
}
