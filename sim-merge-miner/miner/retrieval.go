package miner

import (
	"time"

	"mockbtc/rpc"
	"mockbtc/config"
	"mockbtc/logger"
)

type ToMine struct {
	hash string
	bits uint32 
}

func RetrieveBlocks(config *config.Config) {
	// Mock data for zeroeth retrieval.
	quit := make(chan int)
	toMine := ToMine {
		"e28a262b38316fddefb0b5c753f7cc0022afe94e95f881576ad6b8f33f4e49fa",
		0,
	}
	go MineBlock(config, &toMine, quit)

	// Fetches a block every x second (set in config)
	for true {

		logger.Info.Println("Fetching new block...")

		RetrievedBlock := rpc.RequestCreateAuxBlock(config)

		logger.Info.Println("Babylon Block with hash", RetrievedBlock.Hash ,"received")

		// If hash to mine is changed, terminate current mining efforts and restart
		toMine.bits = RetrievedBlock.Bits

		if RetrievedBlock.Hash != toMine.hash {
			quit <- 0
			toMine.hash = RetrievedBlock.Hash
			quit = make(chan int)
			logger.Info.Println("Prepared BTC block: New mining process starting")
			go MineBlock(config, &toMine, quit)
		}
		time.Sleep(time.Duration(config.Retrieve) * time.Second)
	}
}
