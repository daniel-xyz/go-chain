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
- [ ] > 90% Test Coverage
- [ ] Wallets
- [ ] P2P-Sync

Have fun ❤️ and give feedback!

#### Known issues

- For the purpose of this project, I didn't implement a project structure pattern, like [package oriented design](https://www.ardanlabs.com/blog/2017/02/package-oriented-design.html). The packages are not designed to be used outside of this project.

* It's not considered good practice to pack all types in a common `/types` folder, but it was my first attempt in order to avoid dependency cycles.

#### Start client

Run `make run`

Currently it just runs as a simulation and therefore resets the state with every new run.

It will log the output to stdout, so you'll be seeing what is going on in your console.
Errors will be written into the `errors.log` and the current state of the blockchain persisted in `state.json`. Both files are generated in the root folder.

#### Make API requests

There is a really small API (without validation) just to get the current state and post new transactions from outside the client.

##### GET http://localhost:3001/blockchain

Response will be with `content-type: application/json`

##### POST http://localhost:3001/transaction

Body (`content-type: application/json`):

```json
{
  "from": "any string",
  "to": "any string",
  "amount": "any integer"
}
```

#### Testing

Run `make test`

#### Other useful commands

`make linter`
`make test-coverage-report`
