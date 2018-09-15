package main

import (
	"fmt"

	"github.com/Flur3x/go-chain/block"
)

func main() {
	block := block.New("14-09-2018", "ksh37isdai", "data!")

	fmt.Println(block)
	fmt.Println("Hash is valid: ", block.VerifyHash())
}
