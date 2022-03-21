package main

import (
	"fmt"
	"mockbbld/blockchain"
	"mockbbld/config"
	"mockbbld/logger"
	"mockbbld/pow"
	"mockbbld/rpc"

	"time"

	"github.com/spf13/viper"
)

var i int

func main() {

	//open database
	pow.OpenDB()
	logger.Info.Println("starting to open database")
	defer pow.CloseDB()

	//read config file
	var config config.Config
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	viper.Unmarshal(&config)
	logger.Info.Println("starting to read config file")

	bc := blockchain.NewblockChain(config.Difficulty)

	// start to mine
	go func() {
		t, _ := time.ParseDuration(config.Mtime)
		logger.Info.Println("starting to mine, mining interval:" + config.Mtime)
		for {
			time.Sleep(t) //* time.Second
			i = i + 1
			bc.Addblock([]byte(string(i)))
			logger.Info.Println("generate block, height:", i)
			//mychan1 <- "output1"
		}
	}()

	// start rpc
	rpc.StartRPC(config, bc)

}
