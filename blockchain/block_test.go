package blockchain

import (
	"reflect"
	"testing"

	c "github.com/Flur3x/go-chain/common"
)

func TestNewGenesisBlockReturnsCorrectBlock(t *testing.T) {
	referenceBlock := c.Block{
		Timestamp:    0,
		LastHash:     "-",
		Hash:         "genesis-hash",
		Transactions: []c.Transaction{},
		Difficulty:   6,
		Nonce:        0,
	}

	b := NewGenesisBlock()

	if isEqual := reflect.DeepEqual(referenceBlock, b); !isEqual {
		t.Error("Genesis Block is not correct.")
	}
}
