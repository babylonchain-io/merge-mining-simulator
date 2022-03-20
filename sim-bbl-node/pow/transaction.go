package pow

import (
	"mockbbld/common"
)

type Transaction struct {
	CoinBase string
	Inputs   string
	Outputs  string
	Fee      common.Fixed64
}

func CreateCoinbaseTx(minerAddr string) (Transaction, error) {
	tx := Transaction{
		CoinBase: "CoinBase",
		Inputs:   "test",
		Outputs:  minerAddr,
		Fee:      common.Fixed64(222),
	}
	return tx, nil
}
