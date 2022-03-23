package auxpow

import "mockbtc/common"

//type Fixed64 int64
type AuxBlock struct {
	ChainID           int            `json:"chainid"`
	Height            uint32         `json:"height"`
	CoinBaseValue     common.Fixed64 `json:"coinbasevalue"`
	Bits              uint32         `json:"bits"`
	Hash              string         `json:"hash"`
	PreviousBlockHash string         `json:"previousblockhash"`
}
