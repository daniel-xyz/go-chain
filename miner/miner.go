package miner

import (
	"fmt"

	"github.com/Flur3x/go-chain/blockchain"
	"github.com/Flur3x/go-chain/transactions"
)

// StartMining adds all pending and valid transactions from the transaction pool into a new block.
// This block will the be mined and added to the blockchain state. It repeates this cycle until the client is stopped.
func StartMining(state *blockchain.State) {
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
