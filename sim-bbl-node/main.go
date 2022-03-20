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

//const difficulty = 5

var i int

func main() {

	//open database
	pow.OpenDB()
	defer pow.CloseDB()

	//read coonfig
	var config config.Config
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	viper.Unmarshal(&config)

	bc = blockchain.NewblockChain(config.Difficulty)

	// start to mine
	go func() {
		t, _ := time.ParseDuration(config.Mtime)
		for {
			time.Sleep(t) //* time.Second
			i = i + 1
			bc.Addblock([]byte(string(i)))
			//fmt.Printf("i ---------- %d\n", i)
			//mychan1 <- "output1"
		}
	}()

	// start rpc
	startRPC(config)
}
