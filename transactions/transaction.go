package transactions

import (
	"strings"

	t "github.com/Flur3x/go-chain/types"
	"github.com/google/uuid"
	logging "github.com/op/go-logging"
)

var log = logging.MustGetLogger("")

// New creates a "Transaction" with the given data.
func New(from t.Address, to t.Address, amount uint64, signer t.Signer) (t.Transaction, error) {
	var err error

	outputs := []t.Output{
		t.Output{
			Address: to,
			Amount:  amount,
		},
		t.Output{
			Address: from,
			Amount:  0, // TODO - replace with something like "senderWallet.balance - amount"
		},
	}

	tx := t.Transaction{
		ID: uuid.New(),
		Input: t.Input{
			Address: from,
			Amount:  amount,
		},
		Outputs:   outputs,
		Signature: t.Signature{},
	}

	hash, err := tx.Hash()

	if err != nil {
		return t.Transaction{}, nil
	}

	if tx.Signature, err = signer.Sign(hash); err != nil {
		return t.Transaction{}, err
	}

	return tx, nil
}

// JoinTransactionsToString takes a slice of transactions and returns it as a single string.
func JoinTransactionsToString(txs []t.Transaction) string {
	var stringSlice []string

	for _, tx := range txs {
		stringSlice = append(stringSlice, tx.String())
	}

	return strings.Join(stringSlice, ",")
}
