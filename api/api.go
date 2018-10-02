package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Start initializes the api router with all REST endpoints.
func Start() {
	router := mux.NewRouter()

	router.
		Path("/blockchain").
		HandlerFunc(getBlockchain).
		Methods("GET")

	router.
		Path("/transaction").
		HandlerFunc(postTransaction).
		Methods("POST")

	http.ListenAndServe(":3001", router)
}
