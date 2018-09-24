package blockchain

import (
	"crypto/sha256"
	"fmt"
	"strings"
	"time"

	"github.com/Flur3x/go-chain/transactions"
)

type tx = transactions.Transaction

// Block is chained with other Blocks (lastHash) to form the blockchain.
type Block struct {
	Timestamp    int64
	LastHash     string
	Hash         string
	Transactions []tx
	Difficulty   uint64
	Nonce        int
}

// MineBlock returns a "Block" with given fields and a hash field. Hash is auto-generated based on the given fields.
func MineBlock(lastBlock Block, txs []tx) Block {
	for nonce := 0; ; nonce++ {
		difficulty := currentDifficulty(lastBlock)
		timestamp := currentTimestamp()
		hash := hash(timestamp, lastBlock.Hash, txs, difficulty, nonce)

		if strings.HasPrefix(hash, strings.Repeat("0", int(difficulty))) {
			fmt.Println("Found valid hash: ", hash, "Difficulty: ", difficulty)

			return Block{timestamp, lastBlock.Hash, hash, txs, difficulty, nonce}
		}
	}
}

// NewGenesisBlock returns a "Block" that contains static data. Can be used to initialize a new chain.
func NewGenesisBlock() Block {
	return Block{0, "-", "genesis-hash", []tx{}, 6, 0}
}

// VerifyHash returns "true" if the hash inside the "Block" is a valid hash of itself.
func (b Block) VerifyHash() bool {
	return hash(b.Timestamp, b.LastHash, b.Transactions, b.Difficulty, b.Nonce) == b.Hash
}

func hash(timestamp int64, lastHash string, txs []tx, difficulty uint64, nonce int) string {
	stringifiedIntegers := fmt.Sprintf("%v", timestamp) + fmt.Sprintf("%v", difficulty) + fmt.Sprintf("%v", nonce)
	h := sha256.New()

	h.Write([]byte(lastHash + transactions.JoinTransactionsToString(txs) + stringifiedIntegers))

	return fmt.Sprintf("%x", h.Sum(nil))
}

func currentDifficulty(lastBlock Block) uint64 {
	targetBlockTime := 5 // seconds it should take to mine a block
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

func (b Block) String() string {
	return fmt.Sprintf("::::: Block Info :::::\n\nTimestamp: %d\nLast Hash: %s\nHash: %s\nTransactions: %s\n\n", b.Timestamp, b.LastHash, b.Hash, b.Transactions)
}
