package blockchain

import (
	"reflect"
	"testing"
	"time"

	"github.com/Flur3x/go-chain/types"
	"github.com/stretchr/testify/assert"
)

type fakeTime struct{}

func (f fakeTime) Now() time.Time {
	return time.Date(2018, 10, 9, 12, 48, 30, 0, time.FixedZone("test", 0)) // timestamp: 1539089310
}

func TestNewGenesisBlockReturnsCorrectBlock(t *testing.T) {
	referenceBlock := types.Block{
		Timestamp:    0,
		LastHash:     "-",
		Hash:         "genesis-hash",
		Transactions: []types.Transaction{},
		Difficulty:   6,
		Nonce:        0,
	}

	b := NewGenesisBlock()

	if isEqual := reflect.DeepEqual(referenceBlock, b); !isEqual {
		t.Error("Genesis Block is not correct.")
	}
}

func TestHashReturnsCorrectHash(t *testing.T) {
	b := types.Block{
		Timestamp:    1539075941,
		LastHash:     "0000080800715081b76f2e5dd2fc11ac08900b3d3cb0fb526ffb8dba2c2c7a8d",
		Transactions: []types.Transaction{},
		Difficulty:   4,
		Nonce:        7548,
	}

	h, _ := HashBlockValues(b.Timestamp, b.LastHash, b.Transactions, b.Difficulty, b.Nonce)

	assert.Equal(t, "16b1fde18fd21d4f19d7e009f7e4fb227b239dbae21acc144a045bceec6e7581", h)
}

// func TestMineBlockReturnsBlock(t *testing.T) {
// 	Now = fakeTime{}.Now

// 	expectedBlock := t.Block{
// 		Timestamp: 1539089310,
// 	}

// 	minedBlock, _ := MineBlock([]t.Transaction{})

// 	assert.Equal(t, expectedBlock, minedBlock)
// }
