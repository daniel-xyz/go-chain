package common

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"math/big"

	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
)

// GenerateAddress returns an address that is more readable than the public key, since a public key consists of two values (x and y coordinates).
// https://medium.com/@vedran_s/cryptography-how-to-get-a-bitcoin-address-from-public-key-bd43fed72086
func GenerateAddress(p ecdsa.PublicKey) Address {
	var xy big.Int

	xy.Add(p.X, p.Y)

	// Bitcoin standard to append 1 byte (decimal: 04).
	// Used to distinguish uncompressed public keys from compressed public keys that begin with a 02 or 03.
	pub := "04" + xy.String()

	// Reduce its length by converting it to a byte array and hashing the result.
	op1c := sha256.Sum256([]byte(pub))
	op1 := op1c[:]

	// Improve Security by hashing multiple times.
	op2c := ripemd160.New()
	op2c.Write(op1)
	op2 := op2c.Sum(nil)

	// Add version prefix.
	op3 := append([]byte{0x00}, op2...)

	// Hash again.
	op4c := sha256.Sum256(op3)
	op4 := op4c[:]

	op5c := sha256.Sum256(op4)

	// Only grab first four bytes which are referred to as address checksum.
	op6 := op5c[0:4]

	// Append base of address.
	op7 := append(op3, op6...)

	// Make it more readable before returning it.
	return Address(base58.Encode(op7))
}
