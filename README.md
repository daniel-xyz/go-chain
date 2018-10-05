[![Go Report Card](https://goreportcard.com/badge/github.com/Flur3x/go-chain)](https://goreportcard.com/report/github.com/Flur3x/go-chain) [![GoDoc](https://godoc.org/github.com/Flur3x/go-chain?status.svg)](https://godoc.org/github.com/Flur3x/go-chain)

This is my personal project to get into the Go programming language (Golang) for the first time. You should **not** use this Blockchain or Client, except if you want to have some fun! 🤙

#### Checklist

- [x] Hash & verify blocks
- [x] Add Blocks to Blockchain
- [x] Validate Blockchain
- [x] Real Mining with Dfficulty (nonce)
- [x] Transaction Pool & Add Transactions to Blocks
- [x] Client API
- [x] Error Management
- [ ] Wallets
- [ ] > 90% Test Coverage
- [ ] P2P-Sync

Have fun ❤️ and give feedback!

#### Start client

Run `go run`

It will log the output to stdout, so you'll be seeing what is going on in your console.
Errors will be written into the errors.log and the current state of the blockchain persisted in state.json. Currently it just runs as a simulation and therefore resets the state with every new run.

#### Make API requests

There is a really small API (without validation) just to get the current state and post new transactions from outside the client.

##### GET http://localhost:3001/blockchain

Response will be with `content-type: application/json`

##### POST http://localhost:3001/transaction

Body (`content-type: application/json`):

```json
{
    "from": (address string),
    "to": (address string),
    "amount": (number)
}
```

#### Testing

Run `go test ./..`
