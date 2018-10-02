package miner

import (
	"github.com/Flur3x/go-chain/blockchain"
	"github.com/Flur3x/go-chain/transactions"
)

// Start mining. Collect pending transactions that are valid and add them to a new mined Block.
func Start() error {
	for {
		if err := mine(); err != nil {
			return err
		}
	}
}

func mine() error {
	txs := transactions.ValidTransactions()
	transactions.Clear()

	block, err := blockchain.MineBlock(txs)

	if err != nil {
		return err
	}

	return blockchain.AddBlock(block)
}
