package blockchain

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
)

// Blockchain contains all mined blocks.
type Blockchain struct {
	Blocks []Block
}

// New returns a "Blockchain" struct with the genesis block as the first and only item in it's "Blocks" slice.
func New() Blockchain {
	return Blockchain{Blocks: []Block{NewGenesisBlock()}}
}

// AddBlock adds a "Block" to the given "Blockchain" reference.
func (chain *Blockchain) AddBlock(b Block) {
	chain.Blocks = append(chain.Blocks, b)
}

// IsValidChain validates if all blocks contain valid hashes and are chained properly through their "lastHash". Returns bool.
func (chain *Blockchain) IsValidChain() bool {
	isGenesisBlockValid := cmp.Equal(chain.Blocks[0], NewGenesisBlock())

	hasOnlyValidHashes := func() bool {
		for i := 1; i < len(chain.Blocks); i++ {
			isHashValid := chain.Blocks[i].VerifyHash()
			isLastHashValid := chain.Blocks[i].LastHash == chain.Blocks[i-1].Hash

			if !isHashValid || !isLastHashValid {
				return false
			}
		}

		return true
	}

	return isGenesisBlockValid && hasOnlyValidHashes()
}

func (chain Blockchain) String() string {
	var blocks string

	for _, block := range chain.Blocks {
		blocks = blocks + fmt.Sprintf("%s", block)
	}

	return blocks
}
