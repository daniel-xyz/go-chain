package transactions

import (
	"sync"

	c "github.com/Flur3x/go-chain/common"
	"github.com/google/uuid"
)

var mutex sync.Mutex
var transactionPool = make(map[uuid.UUID]c.Transaction, 100)

// UpdateOrAddToPool takes a Transaction and adds it to the pool.
// If a Transaction with the same ID is already there, it will be replaced.
func UpdateOrAddToPool(transaction c.Transaction) {
	mutex.Lock()

	if _, ok := transactionPool[transaction.ID]; ok {
		defer log.Infof("Transaction updated:\n%s\n", transaction)
	} else {
		defer log.Infof("Transaction created:\n%s\n", transaction)
	}

	transactionPool[transaction.ID] = transaction
	mutex.Unlock()
}

// ValidTransactions returns all transactions within the pool that
// are valid and therefore ready to be added to a Block.
func ValidTransactions() []c.Transaction {
	mutex.Lock()
	validTxs := make([]c.Transaction, 0, len(transactionPool))

	for _, tx := range transactionPool {
		if tx.IsValid() {
			validTxs = append(validTxs, tx)
		}
	}

	mutex.Unlock()

	return validTxs
}

// Clear overrides the transactionPool with a new, empty one.
func Clear() {
	mutex.Lock()
	transactionPool = make(map[uuid.UUID]c.Transaction, 100)
	mutex.Unlock()
}
