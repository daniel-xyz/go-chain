package api

import (
	"encoding/json"
	"net/http"

	"github.com/Flur3x/go-chain/blockchain"
	"github.com/Flur3x/go-chain/transactions"
	t "github.com/Flur3x/go-chain/types"
	"github.com/Flur3x/go-chain/wallet"
)

func getBlockchain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	s, err := blockchain.GetState()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(s); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func postTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var model struct {
		From   t.Address `json:"from"`
		To     t.Address `json:"to"`
		Amount uint64    `json:"amount"`
	}

	myWallet, err := wallet.New()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	t, err := transactions.New(model.From, model.To, model.Amount, myWallet)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&model); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	transactions.UpdateOrAdd(t)
}
