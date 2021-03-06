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

type BBLService struct {
	config config.Config
}

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

	if args.Paytoaddress == "" {
		logger.ReqestError.Println("paytoaddress is nil")
	}

	hash := block.Hash
	hashstring := fmt.Sprintf("%x", hash)

	blockHash := pow.BlockHash{
		Hash: hashstring,
	}
	blockHash.SaveBlockHash()

	//"1d36c855"
	auxBlock := aux.AuxBlock{
		ChainID:           int(249),
		Height:            block.Height + 1,
		CoinBaseValue:     common.Fixed64(175799086),
		Bits:              h.config.Bits,
		Hash:              fmt.Sprintf("%x", hash),
		PreviousBlockHash: fmt.Sprintf("%x", block.Hash),
	}
	*reply = auxBlock
	logger.ReqestInfo.Println("return aux block hash: [" + hashstring + "]")
	return nil
}

func (h *BBLService) SubmitAuxBlock(r *http.Request, args *SubmitAuxArgs, reply *bool) error {

	var aux aux.AuxPow
	blockHashHex := args.Blockhash
	*reply = true

	blocks := bc.GetBlocks()
	block := blocks[len(blocks)-1]
	chash := fmt.Sprintf("%x", block.Hash)

	if chash != blockHashHex {
		*reply = false
		logger.OutdatedError.Println("blockhash is outdated, blockhash: " + blockHashHex)
		return fmt.Errorf("blockhash is outdated, blockhash: " + blockHashHex)
	}

	//block hash check, if blockHashHex is not in our database, return false
	blockhashexit, err := CheckBlockHash(blockHashHex)
	if err != nil {
		*reply = false
		logger.SubmissionError.Println("blockhash not found in database error, blockhash: " + blockHashHex)
		return fmt.Errorf("blockhash not found in database error, blockhash: " + blockHashHex)
	}
	if !blockhashexit {
		*reply = false
		logger.SubmissionError.Println("blockhash not found in database error, blockhash: " + blockHashHex)
		return fmt.Errorf("blockhash not found in database error, blockhash: " + blockHashHex)
	}

	//auxpow deserialization
	auxPow := args.Auxpow
	buf, _ := common.HexStringToBytes(auxPow)
	if err := aux.Deserialize(bytes.NewReader(buf)); err != nil {
		*reply = false
		logger.SubmissionError.Println("auxpow deserialization failed, auxPow: " + auxPow)
		return fmt.Errorf("auxpow deserialization failed, auxPow: " + auxPow)
	}

	bblBits := h.config.Bits
	//auxpow check
	if ok := aux.Check(blockHashHex, 6, bblBits); !ok {
		*reply = false
		logger.SubmissionError.Println("auxpow check failed, auxPow: " + auxPow)
		return fmt.Errorf("auxpow check failed, auxPow: " + auxPow)
	}

	logger.SubmissionInfo.Println("submit auxblock: [" + args.Auxpow + "] from: [" + r.Host + "]")
	return nil
}

func Error(s string) {
	panic("unimplemented")
}

func StartRPC(config config.Config, c *blockchain.Blockchain) {
	bc = c
	logger.Info.Println("starting prc server on port: " + config.Port)
	newServer := rpc.NewServer()
	newServer.RegisterCodec(json.NewCodec(), "application/json")
	bbls := new(BBLService)
	bbls.config = config
	newServer.RegisterService(bbls, "")
	http.Handle("/rpc", newServer)
	http.ListenAndServe(config.Host+":"+config.Port, nil)
}
