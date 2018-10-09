package blockchain

import (
	"reflect"
	"testing"

	"github.com/Flur3x/go-chain/types"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

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

func TestVerifyBlockHashReturnValue(t *testing.T) {
	tests := []struct {
		block          types.Block
		expectedReturn bool
	}{
		{
			block: types.Block{
				Timestamp:    1539092994,
				LastHash:     "0000052effd146db67b85ddee04f2384ae78effe648925e2d4cfbf6559c7a789",
				Hash:         "0000002f1a8f7e1cfa873ae35bf3fa6edaaa7b7b801109d059769a2544ae10d4",
				Transactions: []types.Transaction{},
				Difficulty:   6,
				Nonce:        250240,
			},
			expectedReturn: true,
		},
		{
			block: types.Block{
				Timestamp:    1539092995,
				LastHash:     "0000052effd146db67b85ddee04f2384ae78effa648925e2d4cfbf6559c7a789",
				Hash:         "0000002f1a8f7e1cfa873ae35bf3fa6edaaa7b7b801109d059769a2544ae10d4",
				Transactions: []types.Transaction{},
				Difficulty:   6,
				Nonce:        250240,
			}, expectedReturn: false,
		},
		{
			block: types.Block{
				Timestamp:    1539092994,
				LastHash:     "0000052effd146db67b85ddee04f2384ae78effe648925e2d4cfbf6559c7a789",
				Hash:         "0000002f1a8f7e1cfa873ae35bf3fa6eDaaa7b7b801109d059769a2544ae10d4",
				Transactions: []types.Transaction{},
				Difficulty:   6,
				Nonce:        250240,
			}, expectedReturn: false,
		},
		{
			block: types.Block{
				Timestamp:    1539092994,
				LastHash:     "0000052effd146db67b85ddee04f2384ae78effe648925e2d4cfbf6559c7a789",
				Hash:         "0000002f1a8f7e1cfa873ae35bf3fa6edaaa7b7b801109d059769a2544ae10d4",
				Transactions: []types.Transaction{},
				Difficulty:   7,
				Nonce:        250240,
			}, expectedReturn: false,
		},
		{
			block: types.Block{
				Timestamp: 1539092994,
				LastHash:  "0000052effd146db67b85ddee04f2384ae78effe648925e2d4cfbf6559c7a789",
				Hash:      "0000002f1a8f7e1cfa873ae35bf3fa6edaaa7b7b801109d059769a2544ae10d4",
				Transactions: []types.Transaction{
					{
						ID: uuid.New(),
						Outputs: []types.Output{
							{
								Address: "ozosazuun",
								Amount:  299,
							},
							{
								Address: "blummblumm",
								Amount:  31,
							},
						},
						Input: types.Input{
							Address: "blummblumm",
							Amount:  299,
						},
						Signature: types.Signature{
							R: nil,
							S: nil,
						},
					},
				},
				Difficulty: 6,
				Nonce:      250240,
			},
			expectedReturn: false,
		},
		{
			block: types.Block{
				Timestamp:    1539092994,
				LastHash:     "0000052effd146db67b85ddee04f2384ae78effe648925e2d4cfbf6559c7a789",
				Hash:         "0000002f1a8f7e1cfa873ae35bf3fa6edaaa7b7b801109d059769a2544ae10d4",
				Transactions: []types.Transaction{},
				Difficulty:   6,
				Nonce:        250241,
			},
			expectedReturn: false,
		},
	}

	for _, test := range tests {
		isValid, err := VerifyBlockHash(test.block)

		if assert.NoError(t, err) {
			assert.Equal(t, test.expectedReturn, isValid)
		}
	}
}
