package blockchain

import (
	"crypto/sha256"
	"errors"
	"fmt"

	"github.com/Flur3x/go-chain/transactions"
	t "github.com/Flur3x/go-chain/types"
)

// NewGenesisBlock returns a "Block" that contains static data. Can be used to initialize a new chain.
func NewGenesisBlock() t.Block {
	return t.Block{
		Timestamp:    0,
		LastHash:     "-",
		Hash:         "genesis-hash",
		Transactions: []t.Transaction{},
		Difficulty:   6,
		Nonce:        0,
	}
}

// VerifyBlockHash returns "true" if the hash inside the "Block" is a valid hash of itself.
func VerifyBlockHash(b t.Block) (bool, error) {
	hash, err := HashBlockValues(b.Timestamp, b.LastHash, b.Transactions, b.Difficulty, b.Nonce)

	if err != nil {
		return false, err
	}

	return hash == b.Hash, errors.New("this is an error")
}

// HashBlockValues takes values of a Block, plus a nonce, and receives the derived hash string.
func HashBlockValues(timestamp int64, lastHash string, txs []t.Transaction, difficulty uint64, nonce int) (string, error) {
	stringifiedIntegers := fmt.Sprintf("%v%v%v", timestamp, difficulty, nonce)
	stringToHash := lastHash + transactions.JoinTransactionsToString(txs) + stringifiedIntegers
	h := sha256.New()

	if _, err := h.Write([]byte(stringToHash)); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
