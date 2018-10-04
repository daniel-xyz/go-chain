package common

import (
	"crypto/sha256"
	"fmt"
	"math/big"

	"github.com/google/uuid"
)

// Address of an wallet.
type Address uint64

// Balance of an wallet.
type Balance uint64

// Signature is composed of a pair of integers that is the result of the Sign function.
type Signature struct {
	R *big.Int
	S *big.Int
}

// Input defines how much value an Address receives.
type Input struct {
	Address Address
	Amount  uint64
}

// Output defines how much value an Address sends.
type Output struct {
	Address Address
	Amount  uint64
}

//------------------------------------
// Interfaces
//------------------------------------

// Signer can take a []byte hash and sign it to verify it's validity.
type Signer interface {
	Sign([]byte) (Signature, error)
}

//------------------------------------
// Transaction
//------------------------------------

// Transaction gets created and signed by an account. Will be added to a "Block" by miners.
type Transaction struct {
	ID        uuid.UUID
	Input     Input
	Outputs   []Output
	Signature Signature
}

// IsValid returns "true" if transaction is valid
func (t Transaction) IsValid() bool {
	return t.inputEqualsOutput()
}

func (t Transaction) inputEqualsOutput() bool {
	var totalOutput uint64

	for _, o := range t.Outputs {
		totalOutput += o.Amount
	}

	return totalOutput == t.Input.Amount
}

// Hash returns a sha256 string based on the given transaction.
func (t Transaction) Hash() []byte {
	h := sha256.New()

	h.Write([]byte(fmt.Sprintf("%+v", t)))

	return h.Sum(nil)
}

func (t Transaction) String() string {
	return fmt.Sprintf("\n # ID: %s\n-> Input: %+v\n<- Outputs: %+v\n   Signature: %s\n", t.ID, t.Input, t.Outputs, t.Signature)
}

//------------------------------------
// Block
//------------------------------------

// Block is chained with other Blocks (lastHash) to form the blockchain.
type Block struct {
	Timestamp    int64
	LastHash     string
	Hash         string
	Transactions []Transaction
	Difficulty   uint64
	Nonce        int
}

func (b Block) String() string {
	return fmt.Sprintf("Timestamp: %d\n Last Hash: %s\n Hash: %s\n Difficulty: %d\n Transactions: %s\n\n", b.Timestamp, b.LastHash, b.Hash, b.Difficulty, b.Transactions)
}
