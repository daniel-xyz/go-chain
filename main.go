package main

import (
	"fmt"

	"github.com/Flur3x/go-chain/block"
	"github.com/Flur3x/go-chain/blockchain"
)

func main() {
	chain := blockchain.New()
	b := block.New("14-09-2018", "ksh37isdai", "I'm the second Block!")

	chain.AddBlock(b)
	fmt.Println(chain)
}
