package transactions

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

var mutex sync.Mutex
var transactionPool = make(map[uuid.UUID]Transaction, 100)

// UpdateOrAddToPool takes a Transaction and adds it to the pool.
// If a Transaction with the same ID is already there, it will be replaced.
func UpdateOrAddToPool(transaction Transaction) {
	mutex.Lock()

	if _, ok := transactionPool[transaction.ID]; ok {
		defer fmt.Printf("Transaction updated:\n%s\n", transaction)
	} else {
		defer fmt.Printf("Transaction added:\n%s\n", transaction)
	}

	transactionPool[transaction.ID] = transaction
	mutex.Unlock()
}

// ValidTransactions returns all transactions within the pool that
// are valid and therefore ready to be added to a Block.
func ValidTransactions() []Transaction {
	mutex.Lock()
	validTxs := make([]Transaction, 0, len(transactionPool))

	for _, tx := range transactionPool {
		if tx.isValid() {
			validTxs = append(validTxs, tx)
		}
	}

	mutex.Unlock()

	return validTxs
}

// Clear overrides the transactionPool with a new, empty one.
func Clear() {
	mutex.Lock()
	transactionPool = make(map[uuid.UUID]Transaction, 100)
	mutex.Unlock()
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
