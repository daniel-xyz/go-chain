package blockchain

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
)

// State contains all mined blocks.
type State struct {
	Blocks []Block
}

// New returns a "State" struct with the genesis block as the first and only value in it's "Blocks" slice.
func New() State {
	blockSlice := make([]Block, 1, 100)
	blockSlice[0] = NewGenesisBlock()

	return State{Blocks: blockSlice}
}

// AddBlock adds a "Block" to the given blockchain state.
func (s *State) AddBlock(b Block) {
	s.Blocks = append(s.Blocks, b)
}

// IsValidChain validates if all blocks contain valid hashes and are chained properly through their "lastHash".
func (s *State) IsValidChain() bool {
	isGenesisBlockValid := cmp.Equal(s.Blocks[0], NewGenesisBlock())

	hasOnlyValidHashes := func() bool {
		for i := 1; i < len(s.Blocks); i++ {
			isHashValid := s.Blocks[i].VerifyHash()
			isLastHashValid := s.Blocks[i].LastHash == s.Blocks[i-1].Hash

			if !isHashValid || !isLastHashValid {
				return false
			}
		}

		return true
	}

	return isGenesisBlockValid && hasOnlyValidHashes()
}

// LastBlock returns the last block in the current chain.
func (s State) LastBlock() Block {
	return s.Blocks[len(s.Blocks)-1]
}

func (s State) String() string {
	var blocks string

	for _, block := range s.Blocks {
		blocks = blocks + fmt.Sprintf("%s", block)
	}

	return blocks
}
