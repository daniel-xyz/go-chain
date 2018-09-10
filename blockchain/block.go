package blockchain

import "fmt"

type block struct {
	timestamp string
	lastHash  string
	hash      string
	data      string
}

// GetGenesis returns a block struct that contains static data. Can be used to initialize a new chain.
func GetGenesis() block {
	return block{"10.09.2018 19:48", "-", "0x862b8akn2o7tsb", "some data"}
}

func (b block) String() string {
	return fmt.Sprintf("::::: Block Info :::::\n\nTimestamp: %s\nLast Hash: %s\nHash: %s\nData: %s\n", b.timestamp, b.lastHash, b.hash, b.data)
}
