package wallet

import (
	"crypto/sha256"
	"fmt"

	"github.com/Flur3x/go-chain/common"
)

// Wallet represents a single user/client wallet in the go-chain.
type Wallet struct {
	address    common.Address
	balance    common.Balance
	publicKey  string
	privateKey string
}

// New returns a wallet
func New() Wallet {
	return Wallet{0, 0, "ottopublic", "foobarkey"}
}

// Sign uses the private key of the wallet to cryptographically sign the given hashed value.
func (w Wallet) Sign(v common.Hashable) string {
	h := sha256.New()

	h.Write([]byte(v.Hash() + w.privateKey))

	return fmt.Sprintf("%x", h.Sum(nil))
}
