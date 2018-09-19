package main

import (
	"fmt"

	"github.com/Flur3x/go-chain/blockchain"
	"github.com/Flur3x/go-chain/transactions"
)

func main() {
	testTransactions()
	// testMining()
}

func testMining() {
	chain := blockchain.New()

	block1 := blockchain.MineBlock(chain.Blocks[0], "I'm the second Block!")
	block2 := blockchain.MineBlock(block1, "I'm the third Block!")
	block3 := blockchain.MineBlock(block2, "I'm the third Block!")
	block4 := blockchain.MineBlock(block3, "I'm the third Block!")
	block5 := blockchain.MineBlock(block4, "I'm the third Block!")

	chain.AddBlock(block1)
	chain.AddBlock(block2)
	chain.AddBlock(block3)
	chain.AddBlock(block4)
	chain.AddBlock(block5)

	fmt.Println("\n", chain)
	fmt.Println("Is valid chain: ", chain.IsValidChain())
}

func testTransactions() {
	t := transactions.New(1, 2, 1000)

	transactions.UpdateOrAddToPool(t)
	transactions.UpdateOrAddToPool(t)
}
