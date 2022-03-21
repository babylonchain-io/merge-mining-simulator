package rpc

import (
	"fmt"
	"mockbbld/pow"
)

func CheckBlockHash(blockHashHex string) (bool, error) {
	var blockhashexit bool
	blockhashexit = true

	jsonMap, err := pow.GetBlockHash(blockHashHex)

	if err != nil {
		blockhashexit = false
		fmt.Println(err)
	}

	if jsonMap == nil {
		blockhashexit = false
	}

	return blockhashexit, err
}
