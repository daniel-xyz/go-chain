package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Start initializes the api router with all REST endpoints.
func Start(errorReport chan<- error) {
	router := mux.NewRouter()

	router.
		Path("/blockchain").
		HandlerFunc(getBlockchain).
		Methods("GET")

	router.
		Path("/transaction").
		HandlerFunc(postTransaction).
		Methods("POST")

	if err := http.ListenAndServe(":3001", router); err != nil {
		errorReport <- err
	}
}
