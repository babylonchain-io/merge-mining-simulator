package rpc

import (
	"bytes"
	"fmt"
	"log"
	aux "mockbbld/auxpow"
	"mockbbld/blockchain"
	"mockbbld/common"
	"mockbbld/config"
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
	return nil
}

func (h *BBLService) SubmitAuxBlock(r *http.Request, args *SubmitAuxArgs, reply *bool) error {

	var aux aux.AuxPow
	blockHashHex := args.Blockhash

	//block hash check, if blockHashHex is not in our database, return false
	blockhashexit, err := CheckBlockHash(blockHashHex)
	if err != nil {
		*reply = blockhashexit
		return nil
	}

	//aux pow check
	auxPow := args.Auxpow
	buf, _ := common.HexStringToBytes(auxPow)
	if err := aux.Deserialize(bytes.NewReader(buf)); err != nil {
		*reply = false
		fmt.Printf("auxpow deserialization failed : %s\n", aux)
	}

	blockHash, err := common.Uint256FromHexString(blockHashHex)
	if err != nil {
		fmt.Printf("%s", "bad blockhash")
	}

	if ok := aux.Check(blockHash, 6); !ok {
		*reply = false
		fmt.Printf("auxpow checking failed\n\n\n\n\n\n")
	}
	*reply = true
	return nil
}

func StartRPC(config config.Config, c *blockchain.Blockchain) {
	bc = c
	//fmt.Println("config: ", config, "Host-----: ", config.Host)
	log.Printf("Starting RPC Server on :10000\n")
	newServer := rpc.NewServer()
	newServer.RegisterCodec(json.NewCodec(), "application/json")
	newServer.RegisterService(new(BBLService), "")
	http.Handle("/rpc", newServer)
	//ipport :=
	http.ListenAndServe(config.Host+":"+config.Port, nil)
}
