package main

import (
	"fmt"

	"github.com/Flur3x/go-chain/blockchain"
)

func main() {
	setupTestnet()
}

func setupTestnet() {
	chain := blockchain.New()
	block1 := blockchain.NewBlock("14-09-2018", chain.Blocks[0].Hash, "I'm the second Block!")
	block2 := blockchain.NewBlock("17-09-2018", block1.Hash, "I'm the third Block!")

	chain.AddBlock(block1)
	chain.AddBlock(block2)

	fmt.Println(chain)
	fmt.Println("Is valid chain: ", chain.IsValidChain())
}
