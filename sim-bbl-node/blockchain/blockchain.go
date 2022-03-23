package blockchain

import (
	"fmt"
	"mockbbld/pow"
	pw "mockbbld/pow"
	"time"
)

//Blockchain
const gnnesinfo = "Genesis Block"

var currentHight int32 = 0

//Blockchain structure
type Blockchain struct {
	blocks []*pow.Block
	bits   uint32
}

func newBlock(data []byte, prehash []byte, height uint32, bits uint32) *pow.Block {
	block := pw.Block{
		Version:       00,
		MerkelRoot:    []byte{},
		TimeStamp:     time.Now().Format("2006-15:04:05"),
		Bits:          bits,
		Nonce:         0,
		PrevBlockHash: prehash,
		Data:          data,
		Hash:          prehash,
		Height:        0,
	}

	pow := pw.NewPOW(&block, bits)
	nonce, hash := pow.Mine()
	block.Nonce = nonce
	block.Hash = hash
	block.Height = height

	currentHight = currentHight + 1
	//block.SaveBlock()
	return &block
}

func NewblockChain(bits uint32) *Blockchain {
	var bc Blockchain
	block := newBlock([]byte(gnnesinfo), []byte{}, 0, bits)
	bc.blocks = append(bc.blocks, block)
	bc.bits = bits
	return &bc
}

// add new block
func (this *Blockchain) Addblock(data []byte) {
	lastblockhash := this.blocks[len(this.blocks)-1].Hash
	block := newBlock(data, lastblockhash, uint32(len(this.blocks)-1), this.bits)
	this.blocks = append(this.blocks, block)
	//block.PrintBlockInfo()
}

func (this *Blockchain) GetBlocks() []*pow.Block {
	return this.blocks
}

func (this *Blockchain) PrintAll() {
	for _, v := range this.blocks {
		fmt.Printf("Version : %d\n", v.Version)
		fmt.Printf("PrevBlockHash : %x\n", v.PrevBlockHash)
		fmt.Printf("Hash : %x\n", v.Hash)
		fmt.Printf("MerkleRoot : %x\n", v.MerkelRoot)
		fmt.Printf("TimeStamp : %s\n", v.TimeStamp)
		fmt.Printf("Bits : %d\n", v.Bits)
		fmt.Printf("Nonce : %d\n", v.Nonce)
		fmt.Printf("Data : %s\n", v.Data)
		fmt.Printf("--------------------------------")
	}
}
