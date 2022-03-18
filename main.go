package main

import (
	"fmt"
	"mockbbld/blockchain"
	"mockbbld/config"
	"mockbbld/pow"
	"time"

	"github.com/spf13/viper"
)

var bc *blockchain.Blockchain

const difficulty = 5

var i int

func main() {
	bc = blockchain.NewblockChain(difficulty)

	//open database
	pow.OpenDB()
	defer pow.CloseDB()

	var config config.Config

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	viper.Unmarshal(&config)

	go func() {
		for {
			time.Sleep(30 * time.Second)
			i = i + 1
			bc.Addblock([]byte(string(i)))
			//fmt.Printf("i ---------- %d\n", i)
			//mychan1 <- "output1"
		}
	}()
	// start rpc
	startRPC(config)
}
