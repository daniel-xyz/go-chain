package api

import (
	"encoding/json"
	"net/http"

	"github.com/Flur3x/go-chain/blockchain"

	"github.com/gorilla/mux"
)

var chainState *blockchain.State

func getBlockchain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(chainState)
}

// func postTransaction(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	var t transactions.Transaction

// 	json.NewDecoder(r.Body).Decode(&t)

// 	// transactions.UpdateOrAddToPool(transactions.New(1, 2, uint64(rand.Int63n(10000))))
// }

// Start initializes the api router with all REST endpoints.
func Start(state *blockchain.State) {
	chainState = state
	router := mux.NewRouter()

	router.
		Path("/blockchain").
		HandlerFunc(getBlockchain).
		Methods("GET")

	// router.
	// 	Path("/transaction").
	// 	HandlerFunc(postTransaction).
	// 	Methods("POST")

	http.ListenAndServe(":3001", router)
}
