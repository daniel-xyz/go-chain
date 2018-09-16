package block

import (
	"crypto/sha256"
	"fmt"
)

// Block is chained with other Blocks (lastHash) to form the blockchain.
type Block struct {
	Timestamp string
	LastHash  string
	Hash      string
	Data      string
}

// New returns a "Block" with given fields and a hash field. Hash is auto-generated based on the given fields.
func New(timestamp, lastHash, data string) Block {
	hash := getHash(timestamp, lastHash, data)

	return Block{timestamp, lastHash, hash, data}
}

// NewGenesis returns a "Block" that contains static data. Can be used to initialize a new chain.
func NewGenesis() Block {
	return New("10.09.2018 19:48", "-", "some data 👀")
}

func (b Block) String() string {
	return fmt.Sprintf("::::: Block Info :::::\n\nTimestamp: %s\nLast Hash: %s\nHash: %s\nData: %s\n\n", b.Timestamp, b.LastHash, b.Hash, b.Data)
}

// VerifyHash returns "true" if the hash inside the "Block" is a valid hash of itself.
func (b Block) VerifyHash() bool {
	return getHash(b.Timestamp, b.LastHash, b.Data) == b.Hash
}

func getHash(timestamp, lastHash, data string) string {
	h := sha256.New()

	h.Write([]byte(timestamp + lastHash + data))

	return fmt.Sprintf("%x", h.Sum(nil))
}
