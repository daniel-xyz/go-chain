package api

import (
	"encoding/json"
	"net/http"

	"github.com/Flur3x/go-chain/blockchain"
	c "github.com/Flur3x/go-chain/common"
	"github.com/Flur3x/go-chain/transactions"
	"github.com/Flur3x/go-chain/wallet"
)

func getBlockchain(w http.ResponseWriter, r *http.Request) {
	s, err := blockchain.GetState()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		json.NewEncoder(w).Encode(s)
		w.Header().Set("Content-Type", "application/json")
	}
}

func postTransaction(w http.ResponseWriter, r *http.Request) {
	var t struct {
		From   c.Address `json:"from"`
		To     c.Address `json:"to"`
		Amount uint64    `json:"amount"`
	}

	myWallet := wallet.New()

	json.NewDecoder(r.Body).Decode(&t)
	transactions.UpdateOrAddToPool(transactions.New(t.From, t.To, t.Amount, myWallet))
	w.Header().Set("Content-Type", "application/json")
}
