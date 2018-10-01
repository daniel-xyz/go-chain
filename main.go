package main

import (
	"math/rand"
	"time"

	"github.com/Flur3x/go-chain/api"
	"github.com/Flur3x/go-chain/blockchain"
	"github.com/Flur3x/go-chain/miner"
	"github.com/Flur3x/go-chain/transactions"
)

func main() {
	testRun()
}

func testRun() {
	chain := blockchain.New()

	go api.Start(&chain)
	go miner.Start(&chain)

	for range time.NewTicker(5 * time.Second).C {
		transactions.UpdateOrAddToPool(transactions.New(1, 2, uint64(rand.Int63n(10000))))
	}
}
