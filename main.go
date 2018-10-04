package main

import (
	"math/rand"
	"time"

	"github.com/Flur3x/go-chain/api"
	"github.com/Flur3x/go-chain/blockchain"
	"github.com/Flur3x/go-chain/miner"
	"github.com/Flur3x/go-chain/transactions"
	"github.com/Flur3x/go-chain/wallet"
	logging "github.com/op/go-logging"
)

var log = logging.MustGetLogger("")
var errorReport = make(chan error)

func main() {
	go runSimulation()

	catchErrors()
}

func runSimulation() {
	blockchain.New()

	myWallet, err := wallet.New()

	if err != nil {
		errorReport <- err
	}

	go api.Start(errorReport)
	go miner.Start(errorReport)

	log.Info("\nSimulation started 🌈\n\nFake Transactions are being created and Blocks mined ...\n\n")

	for range time.NewTicker(5 * time.Second).C {
		fakeTransaction, err := transactions.New(1, 2, uint64(rand.Int63n(10000)), myWallet)

		if err != nil {
			errorReport <- err
		}

		transactions.UpdateOrAddToPool(fakeTransaction)
	}
}

func catchErrors() {
	defer func() {
		if r := recover(); r != nil {
			log.Errorf("%+v\n", r)
			log.Infof("More detailed logs in the errors.log file\n")
		}
	}()

	for r := range errorReport {
		panic("Client crashed. Error: " + r.Error())
	}
}
