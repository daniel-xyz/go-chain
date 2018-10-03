package blockchain

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/google/go-cmp/cmp"
	logging "github.com/op/go-logging"
)

var log = logging.MustGetLogger("")

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
func AddBlock(b Block) error {
	s, err := GetState()

	if err != nil {
		return err
	}

	s.Blocks = append(s.Blocks, b)

	if err := setState(s); err != nil {
		return err
	}

	log.Info("::::: Block added to chain :::::\n\n", b)

	return nil
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

func lastBlock() (Block, error) {
	s, err := GetState()

	if err != nil {
		return Block{}, nil
	}

	return s.Blocks[len(s.Blocks)-1], nil
}

// GetState returns the current state of the blockchain with all their blocks.
func GetState() (State, error) {
	return readFromFile()
}

func setState(state State) error {
	return writeToFile(state)
}

func readFromFile() (State, error) {
	var state State

	file, err := stateFile()

	defer file.Close()

	if err != nil {
		return state, err
	}

	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&state); err != nil {
		return state, err
	}

	return state, nil
}

func writeToFile(state State) error {
	file, err := stateFile()

	defer file.Close()

	if err != nil {
		return err
	}

	if err := file.Truncate(0); err != nil {
		return err
	}

	if _, err := file.Seek(0, 0); err != nil {
		return err
	}

	encoder := json.NewEncoder(file)

	if err := encoder.Encode(state); err != nil {
		return err
	}

	return err
}

func stateFile() (*os.File, error) {
	return os.OpenFile("state.json", os.O_RDWR|os.O_CREATE, 0666)
}
