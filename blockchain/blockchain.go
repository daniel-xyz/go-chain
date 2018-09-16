package blockchain

import (
	"github.com/Flur3x/go-chain/block"
)

// Blockchain contains all mined blocks.
type Blockchain struct {
	Blocks []block.Block
}

// New returns a "Blockchain" struct with the genesis block as the first and only item in it's "Blocks" slice.
func New() Blockchain {
	return Blockchain{Blocks: []block.Block{block.NewGenesis()}}
}

// AddBlock adds a "Block" to the given "Blockchain" reference.
func (b *Blockchain) AddBlock(block block.Block) {
	b.Blocks = append(b.Blocks, block)
}
