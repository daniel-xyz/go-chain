package transactions

import (
	"fmt"

	"github.com/google/uuid"
)

var transactionPool = make(map[uuid.UUID]Transaction, 100)

// UpdateOrAddToPool takes a Transaction and adds it to the pool.
// If a Transaction with the same ID is already there, it will be replaced.
func UpdateOrAddToPool(transaction Transaction) {
	if _, ok := transactionPool[transaction.ID]; ok {
		defer fmt.Printf("--- Transaction updated\n%s\n", transaction)
	} else {
		defer fmt.Printf("--- Transaction added\n%s\n", transaction)
	}

	transactionPool[transaction.ID] = transaction
}
