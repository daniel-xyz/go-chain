package miner

import (
	"github.com/Flur3x/go-chain/blockchain"
	"github.com/Flur3x/go-chain/transactions"
)

// Start mining. Collect pending transactions that are valid and add them to a new mined Block.
func Start() {
	for {
		mine()
	}
}

func mine() {
	txs := transactions.ValidTransactions()
	transactions.Clear()

	block := blockchain.MineBlock(txs)
	blockchain.AddBlock(block)
}
