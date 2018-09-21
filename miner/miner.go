package miner

import (
	"github.com/Flur3x/go-chain/blockchain"
	"github.com/Flur3x/go-chain/transactions"
)

// Mine adds all pending and valid transactions from the transaction pool into a new block.
// This block will the be mined and added to the blockchain state.
func Mine(state *blockchain.State) blockchain.Block {
	txs := transactions.ValidTransactions()
	transactions.Clear()

	block := blockchain.MineBlock(state.LastBlock(), txs)
	state.AddBlock(block)

	return block
}
