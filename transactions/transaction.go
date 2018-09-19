package transactions

import "github.com/google/uuid"

// Address of a wallet. Can receive or send value.
type Address uint64

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
	ID     uuid.UUID
	Input  Input
	Output Output
}

// New creates a "Transaction" with the given data.
func New(from Address, to Address, amount uint64) Transaction {
	return Transaction{uuid.New(), Input{to, amount}, Output{from, amount}} // todo replace 1 with uuid
}
