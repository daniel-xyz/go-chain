package main

import (
	"fmt"

	"github.com/Flur3x/go-chain/blockchain"
	"github.com/Flur3x/go-chain/miner"
	"github.com/Flur3x/go-chain/transactions"
)

func main() {
	testRun()
}

func testRun() {
	chain := blockchain.New()

	addTransactions()

	miner.Mine(&chain)
	miner.Mine(&chain)
	miner.Mine(&chain)
	miner.Mine(&chain)
	miner.Mine(&chain)

	fmt.Println("\n", chain)
	fmt.Println("Is valid chain: ", chain.IsValidChain())
}

func addTransactions() {
	transactions.UpdateOrAddToPool(transactions.New(1, 2, 65))
	transactions.UpdateOrAddToPool(transactions.New(3, 2, 210))
	transactions.UpdateOrAddToPool(transactions.New(1, 8, 80))
	transactions.UpdateOrAddToPool(transactions.New(6, 1, 3200))
	transactions.UpdateOrAddToPool(transactions.New(4, 3, 51))
}
