package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"

	c "github.com/Flur3x/go-chain/common"
)

const initialBalance = 5000

// New returns a wallet
func New() (c.Wallet, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	if err != nil {
		return c.Wallet{}, err
	}

	return c.Wallet{
		Address:    c.GenerateAddress(privateKey.PublicKey),
		PublicKey:  privateKey.PublicKey,
		PrivateKey: *privateKey,
		Balance:    initialBalance}, nil
}
