package miner

import (
	"strconv"
	"mockbtc/common"
	aux "mockbtc/auxpow"
	"mockbtc/config"
	"mockbtc/rpc"
	"mockbtc/logger"
)

func MineBlock(config *config.Config, toMine *ToMine, quit chan int) {

	var auxBlockHash *common.Uint256
	
	// Generate mock proof of work - inserting the Babylon blockheader into it
	auxBlockHash, err := common.Uint256FromHexString(toMine.hash)
	if err != nil {
		panic(err)
	}

	auxPow := aux.GenerateAuxPow(*auxBlockHash)

	// Convert bits to target integer
	target := common.CompactToBig(toMine.bits)

	maxNonce := 2147483647

	// Temporary fix
	found := false

	// Change hardcoded to maxNonce
	for nonce := 0; nonce < maxNonce; nonce++ {
		select {
			// Exit when new hash is found
			case <-quit:
				return
			default:
				if !found {
					auxPow.ParBlockHeader.Nonce = uint32(nonce)
					hash := auxPow.ParBlockHeader.Hash()
					hashNum := common.HashToBig(&hash)
					if hashNum.Cmp(target) <= 0 {
						logger.Info.Println("Nonce (" + strconv.Itoa(nonce) + ") found for block hash " + toMine.hash)
						validBlock := rpc.RequestSubmitAuxBlock(config, auxPow, toMine.hash)
						if validBlock {
							logger.Info.Println("Submitted with nonce " + strconv.Itoa(nonce) + "accepted")
						} else {
							logger.Info.Println("Submitted with nonce " + strconv.Itoa(nonce) + " rejected")
						}
						found = true
					}	
				}
		}
	}
}
