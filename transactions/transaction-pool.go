package transactions

import (
	"fmt"

	"github.com/google/uuid"
)

var transactionPool = makeTransactionPoolMap()

// UpdateOrAddToPool takes a Transaction and adds it to the pool.
// If a Transaction with the same ID is already there, it will be replaced.
func UpdateOrAddToPool(transaction Transaction) {
	if _, ok := transactionPool[transaction.ID]; ok {
		defer fmt.Printf("Transaction updated:\n%s\n", transaction)
	} else {
		defer fmt.Printf("Transaction added:\n%s\n", transaction)
	}

	transactionPool[transaction.ID] = transaction
}

// ValidTransactions returns all transactions within the pool that
// are valid and therefore ready to be added to a Block.
func ValidTransactions() []Transaction {
	validTxs := make([]Transaction, 0, len(transactionPool))

	for _, tx := range transactionPool {
		if tx.isValid() {
			validTxs = append(validTxs, tx)
		}
	}

	return validTxs
}

// Clear overrides the transactionPool with a new, empty one.
func Clear() {
	transactionPool = makeTransactionPoolMap()
}

func makeTransactionPoolMap() map[uuid.UUID]Transaction {
	return make(map[uuid.UUID]Transaction, 100)
}

func (t Transaction) isValid() bool {
	return t.inputEqualsOutput()
}

func (t Transaction) inputEqualsOutput() bool {
	var totalOutput uint64

	for _, o := range t.Outputs {
		totalOutput += o.Amount
	}

	return totalOutput == t.Input.Amount
}
