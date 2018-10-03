package blockchain

import (
	"crypto/sha256"
	"fmt"
	"strings"
	"time"

	c "github.com/Flur3x/go-chain/common"
	"github.com/Flur3x/go-chain/transactions"
)

// MineBlock returns a "Block" with given fields and a hash field. Hash is auto-generated based on the given fields.
func MineBlock(txs []c.Transaction) (c.Block, error) {
	lastBlock, err := lastBlock()

	if err != nil {
		return c.Block{}, err
	}

	for nonce := 0; ; nonce++ {
		difficulty := currentDifficulty(lastBlock)
		timestamp := currentTimestamp()
		hash := hash(timestamp, lastBlock.Hash, txs, difficulty, nonce)

		if strings.HasPrefix(hash, strings.Repeat("0", int(difficulty))) {
			return c.Block{
				Timestamp:    timestamp,
				LastHash:     lastBlock.Hash,
				Hash:         hash,
				Transactions: txs,
				Difficulty:   difficulty,
				Nonce:        nonce,
			}, nil
		}
	}
}

// NewGenesisBlock returns a "Block" that contains static data. Can be used to initialize a new chain.
func NewGenesisBlock() c.Block {
	return c.Block{
		Timestamp:    0,
		LastHash:     "-",
		Hash:         "genesis-hash",
		Transactions: []c.Transaction{},
		Difficulty:   6,
		Nonce:        0,
	}
}

// VerifyHash returns "true" if the hash inside the "Block" is a valid hash of itself.
func VerifyHash(b c.Block) bool {
	return hash(b.Timestamp, b.LastHash, b.Transactions, b.Difficulty, b.Nonce) == b.Hash
}

func hash(timestamp int64, lastHash string, txs []c.Transaction, difficulty uint64, nonce int) string {
	stringifiedIntegers := fmt.Sprintf("%v%v%v", timestamp, difficulty, nonce)
	h := sha256.New()

	h.Write([]byte(lastHash + transactions.JoinTransactionsToString(txs) + stringifiedIntegers))

	return fmt.Sprintf("%x", h.Sum(nil))
}

func currentDifficulty(lastBlock c.Block) uint64 {
	targetBlockTime := 15 // seconds it should take to mine a block
	currentTimestamp := currentTimestamp()
	increaseDifficulty := lastBlock.Timestamp+int64(targetBlockTime) > currentTimestamp

	if increaseDifficulty {
		return lastBlock.Difficulty + 1
	}

	if lastBlock.Difficulty > 1 {
		return lastBlock.Difficulty - 1
	}

	return 1
}

func currentTimestamp() int64 {
	return time.Now().Unix()
}
