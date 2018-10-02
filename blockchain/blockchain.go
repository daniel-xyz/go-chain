package blockchain

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/google/go-cmp/cmp"
)

// State contains the general blockchain state, most importantly all mined blocks.
type State struct {
	Blocks []Block
}

// New returns a "State" struct with the genesis block as the first and only value in it's "Blocks" slice.
func New() State {
	blockSlice := make([]Block, 1, 100)
	blockSlice[0] = NewGenesisBlock()

	state := State{Blocks: blockSlice}

	setState(state)

	return state
}

// AddBlock adds a "Block" to the given blockchain state.
func AddBlock(b Block) {
	s := GetState()
	s.Blocks = append(s.Blocks, b)

	setState(s)

	fmt.Println("Added Block to chain: \n\n", b)
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

func (s State) String() string {
	var blocks string

	for _, block := range s.Blocks {
		blocks = blocks + fmt.Sprintf("%s", block)
	}

	return blocks
}

func lastBlock() Block {
	s := GetState()

	return s.Blocks[len(s.Blocks)-1]
}

// GetState returns the current state of the blockchain with all their blocks.
func GetState() State {
	return readFromFile()
}

func setState(state State) error {
	return writeToFile(state)
}

func readFromFile() State {
	file, err := stateFile()

	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	var state State

	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&state); err != nil {
		log.Fatal(err)
	}

	return state
}

func writeToFile(state State) error {
	file, err := stateFile()

	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	encoder := json.NewEncoder(file)

	if err := encoder.Encode(state); err != nil {
		log.Fatal(err)
	}

	return err
}

func stateFile() (*os.File, error) {
	return os.OpenFile("state.json", os.O_RDWR|os.O_CREATE, 0666)
}
