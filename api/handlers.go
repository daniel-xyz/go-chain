package api

import (
	"encoding/json"
	"net/http"

	"github.com/Flur3x/go-chain/wallet"

	"github.com/Flur3x/go-chain/blockchain"
	"github.com/Flur3x/go-chain/transactions"
)

func getBlockchain(w http.ResponseWriter, r *http.Request) {
	s := blockchain.GetState()

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(s)
}

func postTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var t struct {
		From   wallet.Address `json:"from"`
		To     wallet.Address `json:"to"`
		Amount uint64         `json:"amount"`
	}

	json.NewDecoder(r.Body).Decode(&t)

	transactions.UpdateOrAddToPool(transactions.New(t.From, t.To, t.Amount))
}
