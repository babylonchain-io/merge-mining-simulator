package main

import (
	"mockbtc/miner"
	"mockbtc/config"
	"mockbtc/logger"

	"github.com/spf13/viper"
)

func main() {
	var config config.Config
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	
	if err != nil {
		logger.Info.Println(err)
		return
	}
	viper.Unmarshal(&config)
	
	// Periodically retrieve nodes
	miner.RetrieveBlocks(&config)
}
