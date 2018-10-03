package common

import (
	"crypto/sha256"
	"fmt"

	"github.com/google/uuid"
)

// Address of an wallet.
type Address uint64

// Balance of an wallet.
type Balance uint64

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

// Hashable can be cryptographically hashed. Should return a hash string represantation of the pointer.
type Hashable interface {
	Hash() string
}

// Signer can take a value and sign it to verify it's validity.
type Signer interface {
	Sign(Hashable) string
}

//------------------------------------
// Transaction
//------------------------------------

// Transaction gets created and signed by an account. Will be added to a "Block" by miners.
type Transaction struct {
	ID        uuid.UUID
	Input     Input
	Outputs   []Output
	Signature string
}

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
func (t Transaction) Hash() string {
	h := sha256.New()

	h.Write([]byte(fmt.Sprintf("%+v", t)))

	return fmt.Sprintf("%x", h.Sum(nil))
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
