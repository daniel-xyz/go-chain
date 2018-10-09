package types

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"

	"github.com/google/uuid"
)

// Address of an wallet that is derived from the publicKey.
type Address string

// Balance of an wallet.
type Balance uint64

// Signature is composed of a pair of integers that is the result of the Sign function.
type Signature struct {
	R *big.Int `json:"r"`
	S *big.Int `json:"s"`
}

// Input defines how much value an address receives.
type Input struct {
	Address Address `json:"address"`
	Amount  uint64  `json:"amount"`
}

// Output defines how much value an address sends.
type Output struct {
	Address Address `json:"address"`
	Amount  uint64  `json:"amount"`
}

//------------------------------------
// Interfaces
//------------------------------------

// Signer can take a []byte hash and sign it to verify it's validity.
type Signer interface {
	Sign([]byte) (Signature, error)
}

//------------------------------------
// Wallet
//------------------------------------

// Wallet represents a single user/client wallet in the go-chain.
type Wallet struct {
	Address    Address          `json:"address"`
	PublicKey  ecdsa.PublicKey  `json:"publicKey"`
	PrivateKey ecdsa.PrivateKey `json:"privateKey"`
	Balance    Balance          `json:"balance"`
}

// Sign uses the private key of the wallet to cryptographically sign the given hashed value.
func (w Wallet) Sign(h []byte) (Signature, error) {
	r, s, err := ecdsa.Sign(rand.Reader, &w.PrivateKey, h[:])

	if err != nil {
		return Signature{}, err
	}

	return Signature{R: r, S: s}, nil
}

func (w Wallet) String() string {
	return fmt.Sprintf("\n Address: %0x\n Private Key: %s\n Balanace %d\n", w.Address, w.PrivateKey.D, w.Balance)
}

//------------------------------------
// Transaction
//------------------------------------

// Transaction gets created and signed by an account. Will be added to a "Block" by miners.
type Transaction struct {
	ID        uuid.UUID `json:"id"`
	Input     Input     `json:"input"`
	Outputs   []Output  `json:"outputs"`
	Signature Signature `json:"signature"`
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
func (t Transaction) Hash() ([]byte, error) {
	h := sha256.New()

	if _, err := h.Write([]byte(fmt.Sprintf("%+v", t))); err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

func (t Transaction) String() string {
	return fmt.Sprintf("\n # ID: %s\n-> Input: %+v\n<- Outputs: %+v\n   Signature: %s\n", t.ID, t.Input, t.Outputs, t.Signature)
}

//------------------------------------
// Block
//------------------------------------

// Block is chained with other Blocks (lastHash) to form the blockchain.
type Block struct {
	Timestamp    int64         `json:"timestamp"`
	LastHash     string        `json:"lastHash"`
	Hash         string        `json:"hash"`
	Transactions []Transaction `json:"transactions"`
	Difficulty   uint64        `json:"difficulty"`
	Nonce        int           `json:"nonce"`
}

func (b Block) String() string {
	return fmt.Sprintf("Timestamp: %d\n Last Hash: %s\n Hash: %s\n Difficulty: %d\n Transactions: %s\n\n", b.Timestamp, b.LastHash, b.Hash, b.Difficulty, b.Transactions)
}
