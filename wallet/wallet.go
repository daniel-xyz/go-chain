package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"

	"github.com/Flur3x/go-chain/crypto"
	t "github.com/Flur3x/go-chain/types"
)

const initialBalance = 5000

// New returns a wallet
func New() (t.Wallet, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	if err != nil {
		return t.Wallet{}, err
	}

	return t.Wallet{
		Address:    crypto.GenerateAddress(privateKey.PublicKey),
		PublicKey:  privateKey.PublicKey,
		PrivateKey: *privateKey,
		Balance:    initialBalance}, nil
}
