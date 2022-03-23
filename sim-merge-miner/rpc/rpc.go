package rpc


import (
	"bytes"
	"io/ioutil"

	"encoding/json"
	"net/http"

	aux "mockbtc/auxpow"
	"mockbtc/config"
	"mockbtc/common"
)

type CreateAuxBlockParams struct {
	Paytoaddress string `json:"paytoaddress"`
}

type CreateAuxBlockArgs struct {
	ID string `json:"id"`
	Method string `json:"method"`
	Params[] CreateAuxBlockParams `json:"params"`
}

type SubmitAuxBlockParams struct {
	Blockhash string `json:"blockhash"`
	Auxpow string `json:"auxpow"`
}

type SubmitAuxBlockArgs struct {
	ID string `json:"id"`
	Method string `json:"method"`
	Params[] SubmitAuxBlockParams `json:"params"`
}

func getEndpoint(config *config.Config) string {
	getEndpoint := config.Bblhost + ":" + config.Bblport + "/rpc"
	return getEndpoint
}

func RequestAndUnpack(config *config.Config, requestBody []byte) map[string]interface{} {
	resp, err := http.Post("http://127.0.0.1:10000/rpc", "application/json", bytes.NewBuffer(requestBody))

	if err != nil { panic(err) }

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil { panic(err) }

	var obj map[string]interface{}
	err = json.Unmarshal(body, &obj)

	return obj
}

func RequestCreateAuxBlock(config *config.Config) aux.AuxBlock {
	args := CreateAuxBlockArgs {
		ID: "1",
		Method: "BBLService.Createauxblock",
		Params: []CreateAuxBlockParams{
			CreateAuxBlockParams {
				Paytoaddress: config.Paytoaddress,
			},
		},
	}

	requestBody, err := json.Marshal(args)

	if err != nil { panic(err) }
		
	obj := RequestAndUnpack(config, requestBody)
	result := obj["result"].(map[string]interface{})

	RetrievedBlock := aux.AuxBlock {
		ChainID: int(result["chainid"].(float64)),
		Height: uint32(result["height"].(float64)),
		CoinBaseValue: common.Fixed64(result["coinbasevalue"].(float64)),
		Bits: uint32(result["bits"].(float64)),
		Hash: result["hash"].(string),
		PreviousBlockHash: result["previousblockhash"].(string),
	}

	return RetrievedBlock
}

func RequestSubmitAuxBlock(config *config.Config, auxPow *aux.AuxPow, blockHash string) bool {
	buf := new(bytes.Buffer)
	auxPow.Serialize(buf)
	auxPowSer := common.BytesToHexString(buf.Bytes())

	args := SubmitAuxBlockArgs{
		ID: "1",
		Method: "BBLService.SubmitAuxBlock",
		Params: []SubmitAuxBlockParams{
			SubmitAuxBlockParams {
				Blockhash: blockHash,
				Auxpow: auxPowSer,
			},
		},
	}

	requestBody, err := json.Marshal(args)

	if err != nil { panic(err) }

	obj := RequestAndUnpack(config, requestBody)

	return obj["result"].(bool)
}
