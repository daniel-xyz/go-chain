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
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s)

}

func postTransaction(w http.ResponseWriter, r *http.Request) {
	var model struct {
		From   c.Address `json:"from"`
		To     c.Address `json:"to"`
		Amount uint64    `json:"amount"`
	}

	myWallet, err := wallet.New()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	t, err := transactions.New(model.From, model.To, model.Amount, myWallet)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&model)
	transactions.UpdateOrAddToPool(t)
}
