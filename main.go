package main

import (
	"mockbbld/blockchain"
	"mockbbld/pow"
	"time"
)

var bc *blockchain.Blockchain

const difficulty = 5

var i int

func main() {
	bc = blockchain.NewblockChain(difficulty)

	//open database
	pow.OpenDB()
	defer pow.CloseDB()

	go func() {
		for {
			time.Sleep(5 * time.Second)
			i = i + 1
			bc.Addblock([]byte(string(i)))
			//fmt.Printf("i ---------- %d\n", i)
			//mychan1 <- "output1"
		}
	}()
	// start rpc
	startRPC()
}
