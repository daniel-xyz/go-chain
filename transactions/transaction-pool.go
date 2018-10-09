package transactions

import (
	"sync"

	t "github.com/Flur3x/go-chain/types"
	"github.com/google/uuid"
)

var mutex sync.Mutex
var transactionPool = make(map[uuid.UUID]t.Transaction, 100)

// UpdateOrAdd takes a Transaction and adds it to the pool.
// If a Transaction with the same ID is already there, it will be replaced.
func UpdateOrAdd(transaction t.Transaction) {
	mutex.Lock()
	defer mutex.Unlock()

	if _, ok := transactionPool[transaction.ID]; ok {
		defer log.Infof("Transaction updated:\n%s\n", transaction)
	} else {
		defer log.Infof("Transaction created:\n%s\n", transaction)
	}

	transactionPool[transaction.ID] = transaction
}

// PopValidTransactions returns all transactions within the pool that
// are valid and therefore ready to be added to a Block.
func PopValidTransactions() []t.Transaction {
	mutex.Lock()
	defer clearTransactionPool()
	defer mutex.Unlock()

	validTxs := make([]t.Transaction, 0, len(transactionPool))

	for _, tx := range transactionPool {
		if tx.IsValid() {
			validTxs = append(validTxs, tx)
		}
	}

	return validTxs
}

func clearTransactionPool() {
	mutex.Lock()
	defer mutex.Unlock()

	transactionPool = make(map[uuid.UUID]t.Transaction, 100)
}
