package main

import (
	"fmt"

	"github.com/Flur3x/go-chain/blockchain"
)

func main() {
	genesisBlock := blockchain.GetGenesis()

	fmt.Printf("%+v", genesisBlock)
}
