package rpc

import (
	"bytes"
	"fmt"
	aux "mockbbld/auxpow"
	"mockbbld/blockchain"
	"mockbbld/common"
	"mockbbld/config"
	"mockbbld/logger"
	"mockbbld/pow"
	"net/http"

	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json"
)

var bc *blockchain.Blockchain

type BBLService struct{}

type CreateAuxHashArgs struct {
	Paytoaddress string
}

type SubmitAuxArgs struct {
	Blockhash string
	Auxpow    string
}

func (h *BBLService) Createauxblock(r *http.Request, args *CreateAuxHashArgs, reply *aux.AuxBlock) error {

	//var baseTx pow.Transaction
	blocks := bc.GetBlocks()
	block := blocks[len(blocks)-1]
	//baseTx, _ = pow.CreateCoinbaseTx(args.Paytoaddress)
	//bc.Addblock([]byte(fmt.Sprintf("%v", baseTx))) //add baseTx to block
	//new_pow := pow.NewPOW(block, difficulty)
	//hash := new_pow.PreparetoMine()
	hash := block.Hash
	hashstring := fmt.Sprintf("%x", hash)

	blockHash := pow.BlockHash{
		Hash: hashstring,
	}
	blockHash.SaveBlockHash()

	auxBlock := aux.AuxBlock{
		ChainID:           int(249),
		Height:            block.Height + 1,
		CoinBaseValue:     common.Fixed64(175799086),
		Bits:              "1d36c855",
		Hash:              fmt.Sprintf("%x", hash),
		PreviousBlockHash: fmt.Sprintf("%x", block.Hash),
	}
	*reply = auxBlock
	logger.Info.Println("return aux block hash: [" + hashstring + "] to: [" + r.Host + "]")
	return nil
}

func (h *BBLService) SubmitAuxBlock(r *http.Request, args *SubmitAuxArgs, reply *bool) error {

	var aux aux.AuxPow
	blockHashHex := args.Blockhash
	*reply = true

	//block hash check, if blockHashHex is not in our database, return false
	/*
		blockhashexit, err := CheckBlockHash(blockHashHex)
		if err != nil {
			*reply = false
			logger.Error.Println("block database error")
			return nil
		}
		if !blockhashexit {
			*reply = false
			logger.Error.Println("not found block hash")
			return nil
		}
	*/

	//auxpow deserialization
	auxPow := args.Auxpow
	buf, _ := common.HexStringToBytes(auxPow)
	if err := aux.Deserialize(bytes.NewReader(buf)); err != nil {
		*reply = false
		logger.Error.Println("auxpow deserialization failed")
		return nil
	}

	//auxpow check
	if ok := aux.Check(blockHashHex, 8); !ok {
		*reply = false
		logger.Error.Println("auxpow checking failed")
		return nil
	}
	return nil
}

func StartRPC(config config.Config, c *blockchain.Blockchain) {
	bc = c
	logger.Info.Println("starting prc server on port: " + config.Port)
	newServer := rpc.NewServer()
	newServer.RegisterCodec(json.NewCodec(), "application/json")
	newServer.RegisterService(new(BBLService), "")
	http.Handle("/rpc", newServer)
	http.ListenAndServe(config.Host+":"+config.Port, nil)
}
