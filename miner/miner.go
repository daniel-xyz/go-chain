package miner

import (
	"fmt"

	"github.com/Flur3x/go-chain/blockchain"
	"github.com/Flur3x/go-chain/transactions"
)

// Start mining. Collect pending transactions that are valid and add them to a new mined Block.
func Start(state *blockchain.State) {
	for {
		mine(state)
	}
}

func mine(state *blockchain.State) blockchain.Block {
	txs := transactions.ValidTransactions()
	transactions.Clear()

	block := blockchain.MineBlock(state.LastBlock(), txs)
	state.AddBlock(block)

	fmt.Println("Added Block to chain: \n\n", block)

	return block
}
