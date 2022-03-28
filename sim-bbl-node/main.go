package main

import (
	"fmt"
	"mockbbld/blockchain"
	"mockbbld/config"
	"mockbbld/logger"
	"mockbbld/pow"
	"mockbbld/rpc"

	"bufio"
	"os"
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

	bc := blockchain.NewblockChain(config.Bits)

	// start to mine
	go func() {
		t, _ := time.ParseDuration(config.Mtime)
		logger.Info.Println("starting to mine, mining interval:" + config.Mtime)
		for {
			time.Sleep(t) //* time.Second
			i = i + 1
			bc.Addblock([]byte(string(i)))
			logger.BlockInfo.Println("generate block, height:", i)
			//mychan1 <- "output1"
		}
	}()

	// start rpc
	go func() {
		rpc.StartRPC(config, bc)
	}()

	// read command
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if line == "exit" {
			break
		}

		if line == "show_normal_aux_request" {
			logger.ShowNormalAuxRequest()
		}

		if line == "show_normal_aux_submission" {
			logger.ShowNormalAuxSubmission()
		}

		if line == "show_error_aux_request" {
			logger.ShowErrorAuxRequest()
		}

		if line == "show_error_aux_submission" {
			logger.ShowErrorAuxSubmission()
		}

		if line == "show_all" {
			logger.ShowAll()
		}

	}

}
