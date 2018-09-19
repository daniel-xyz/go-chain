package transactions

import (
	"fmt"

	"github.com/Flur3x/go-chain/wallet"
	"github.com/google/uuid"
)

// Address of an account. Can receive or send value.
type Address = wallet.Address

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

// Transaction gets created and signed by an account. Will be added to a "Block" by miners.
type Transaction struct {
	ID      uuid.UUID
	Input   Input
	Outputs []Output
}

// New creates a "Transaction" with the given data.
func New(from Address, to Address, amount uint64) Transaction {
	outputs := []Output{
		Output{to, amount},
		Output{from, 0}, // todo replace with something like "senderWallet.balance - amount"
	}

	return Transaction{uuid.New(), Input{to, amount}, outputs}
}

func (t Transaction) String() string {
	return fmt.Sprintf("ID: %s\nInput: %+v\nOutputs: %+v\n", t.ID, t.Input, t.Outputs)
}
