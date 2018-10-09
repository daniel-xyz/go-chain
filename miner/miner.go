package miner

import (
	"strings"
	"time"

	"github.com/Flur3x/go-chain/blockchain"
	"github.com/Flur3x/go-chain/transactions"
	t "github.com/Flur3x/go-chain/types"
)

var (
	// Now can be overriden by tests
	Now = time.Now
)

// Start mining. Collect pending transactions that are valid and add them to a new mined Block.
func Start(errorReport chan<- error) {
	for {
		if err := mineBlock(); err != nil {
			errorReport <- err
		}
	}
}

func mineBlock() error {
	txs := transactions.PopValidTransactions()
	block, err := findValidNonce(txs)

	if err != nil {
		return err
	}

	return blockchain.AddBlock(block)
}

func findValidNonce(txs []t.Transaction) (t.Block, error) {
	for nonce := 0; ; nonce++ {
		block, isValid, err := tryNonce(nonce, txs)

		if err != nil {
			return t.Block{}, err
		}

		if isValid {
			return block, nil
		}
	}
}

func tryNonce(nonce int, txs []t.Transaction) (t.Block, bool, error) {
	lastBlock, err := blockchain.LastBlock()

	if err != nil {
		return t.Block{}, false, nil
	}

	difficulty := currentDifficulty(lastBlock)
	timestamp := currentTimestamp()
	hash, err := blockchain.HashBlockValues(timestamp, lastBlock.Hash, txs, difficulty, nonce)

	if err != nil {
		return t.Block{}, false, nil
	}

	hasFoundValidHash := strings.HasPrefix(hash, strings.Repeat("0", int(difficulty)))

	if hasFoundValidHash {
		return t.Block{
			Timestamp:    timestamp,
			LastHash:     lastBlock.Hash,
			Hash:         hash,
			Transactions: txs,
			Difficulty:   difficulty,
			Nonce:        nonce,
		}, true, nil
	}

	return t.Block{}, false, nil
}

func currentDifficulty(lastBlock t.Block) uint64 {
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
	return Now().Unix()
}
