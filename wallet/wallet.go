package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"

	c "github.com/Flur3x/go-chain/common"
)

const initialBalance = 5000

// Wallet represents a single user/client wallet in the go-chain.
type Wallet struct {
	address    c.Address
	balance    c.Balance
	publicKey  ecdsa.PublicKey
	privateKey ecdsa.PrivateKey
}

// New returns a wallet
func New() (Wallet, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	if err != nil {
		return Wallet{}, err
	}

	return Wallet{0, initialBalance, privateKey.PublicKey, *privateKey}, nil
}

// Sign uses the private key of the wallet to cryptographically sign the given hashed value.
func (w Wallet) Sign(h []byte) (c.Signature, error) {
	r, s, err := ecdsa.Sign(rand.Reader, &w.privateKey, h[:])

	if err != nil {
		return c.Signature{}, err
	}

	return c.Signature{R: r, S: s}, nil
}
