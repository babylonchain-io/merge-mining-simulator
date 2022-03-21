package rpc

import (
	"mockbbld/pow"
)

func CheckBlockHash(blockHashHex string) (bool, error) {
	var blockhashexit bool
	blockhashexit = true

	jsonMap, err := pow.GetBlockHash(blockHashHex)

	if err != nil {
		blockhashexit = false
	}

	if jsonMap == nil {
		blockhashexit = false
	}
	return blockhashexit, err
}
