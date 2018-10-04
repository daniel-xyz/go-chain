package transactions

import (
	"strings"

	c "github.com/Flur3x/go-chain/common"
	"github.com/google/uuid"
	logging "github.com/op/go-logging"
)

var log = logging.MustGetLogger("")

// New creates a "Transaction" with the given data.
func New(from c.Address, to c.Address, amount uint64, signer c.Signer) (c.Transaction, error) {
	var err error

	outputs := []c.Output{
		c.Output{Address: to, Amount: amount},
		c.Output{Address: from, Amount: 0}, // TODO - replace with something like "senderWallet.balance - amount"
	}

	t := c.Transaction{
		ID: uuid.New(),
		Input: c.Input{
			Address: from,
			Amount:  amount,
		},
		Outputs:   outputs,
		Signature: c.Signature{},
	}

	if t.Signature, err = signer.Sign(t.Hash()); err != nil {
		return c.Transaction{}, err
	}

	return t, nil
}

// JoinTransactionsToString takes a slice of transactions and returns it as a single string.
func JoinTransactionsToString(txs []c.Transaction) string {
	var stringSlice []string

	for _, tx := range txs {
		stringSlice = append(stringSlice, tx.String())
	}

	return strings.Join(stringSlice, ",")
}
