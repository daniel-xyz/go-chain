package transactions

import (
	"fmt"

	"github.com/google/uuid"
)

var transactionPool = map[uuid.UUID]Transaction{}

// UpdateOrAddToPool takes a Transaction and adds it to the pool.
// If a Transaction with the same ID is already there, it will be replaced.
func UpdateOrAddToPool(transaction Transaction) {
	if _, ok := transactionPool[transaction.ID]; ok {
		defer fmt.Println("Transaction updated: ", transaction)
	} else {
		fmt.Println("Transaction added: ", transaction)
	}

	transactionPool[transaction.ID] = transaction
}
