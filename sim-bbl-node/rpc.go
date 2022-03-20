package main

import (
	"bytes"
	"fmt"
	"log"
	aux "mockbbld/auxpow"
	"mockbbld/common"
	"mockbbld/config"
	"mockbbld/pow"
	"net/http"

	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json"
)

type Fixed64 int64
type BBLService struct{}

type CreateAuxHashArgs struct {
	Paytoaddress string
}

type AuxBlock struct {
	ChainID           int     `json:"chainid"`
	Height            uint32  `json:"height"`
	CoinBaseValue     Fixed64 `json:"coinbasevalue"`
	Bits              string  `json:"bits"`
	Hash              string  `json:"hash"`
	PreviousBlockHash string  `json:"previousblockhash"`
}

type SubmitAuxArgs struct {
	Blockhash string
	Auxpow    string
}

func (h *BBLService) Createauxblock(r *http.Request, args *CreateAuxHashArgs, reply *AuxBlock) error {

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

	//height, err := strconv.ParseUint(block.Height, 10, 32)
	//if err != nil {
	//	fmt.Println(err)
	//}

	auxBlock := AuxBlock{
		ChainID:           int(249),
		Height:            block.Height + 1,
		CoinBaseValue:     Fixed64(175799086),
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

	blockHash, err := common.Uint256FromHexString(blockHashHex)
	if err != nil {
		fmt.Printf("%s", "bad blockhash")
	}
	jsonMap, err := pow.GetBlockHash(blockHashHex)

	if err != nil {
		*reply = false
		fmt.Println(err)
	}

	if jsonMap == nil {
		*reply = false
		fmt.Println(jsonMap)
	}

	//aux pow check
	auxPow := args.Auxpow

	buf, _ := common.HexStringToBytes(auxPow)
	if err := aux.Deserialize(bytes.NewReader(buf)); err != nil {
		fmt.Printf("auxpow deserialization failed : %s\n", aux)
	}

	if ok := aux.Check(blockHash, 6); !ok {
		fmt.Printf("auxpow checking failed\n\n\n\n\n\n")
	}

	//fmt.Printf("Auxpow : %s\n", auxPow)

	*reply = true
	return nil
}

func startRPC(config config.Config) {
	//fmt.Println("config: ", config, "Host-----: ", config.Host)
	log.Printf("Starting RPC Server on :10000\n")
	newServer := rpc.NewServer()
	newServer.RegisterCodec(json.NewCodec(), "application/json")
	newServer.RegisterService(new(BBLService), "")
	http.Handle("/rpc", newServer)
	//ipport :=
	http.ListenAndServe(config.Host+":"+config.Port, nil)
}
