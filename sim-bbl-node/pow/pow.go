package pow

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math/big"
)

type Block struct {
	Version       uint64
	MerkelRoot    []byte
	TimeStamp     string
	Difficulty    uint64
	Nonce         uint64
	PrevBlockHash []byte
	Data          []byte
	Hash          []byte
	Height        uint32
}

type BlockHash struct {
	Hash string
}

func NewPOW(block *Block, difficulty uint64) *ProofOfWork {
	var this ProofOfWork
	this.block = block
	targetint := Gettargetint(difficulty)
	this.target = targetint
	return &this
}

func (block *Block) PrintBlockInfo() {

	fmt.Printf("=========Block Hight %d=========\n", block.Height)
	fmt.Printf("Version : %d\n", block.Version)
	fmt.Printf("PrevBlockHash : %x\n", block.PrevBlockHash)
	fmt.Printf("Hash : %x\n", block.Hash)
	fmt.Printf("MerkleRoot : %x\n", block.MerkelRoot)
	fmt.Printf("TimeStamp : %s\n", block.TimeStamp)
	fmt.Printf("Difficuty : %d\n", block.Difficulty)
	fmt.Printf("Nonce : %d\n", block.Nonce)
	fmt.Printf("Data : %s\n", block.Data)
	fmt.Printf("--------------------------------\n\n")
}

type ProofOfWork struct {
	target *big.Int
	block  *Block
}

func Gettargetint(difficulty uint64) *big.Int {
	targetint := big.NewInt(1)
	targetint.Lsh(targetint, uint(256-difficulty))
	return targetint
}

func uint2byte(num uint64) []byte {
	var buff bytes.Buffer
	binary.Write(&buff, binary.BigEndian, &num)
	return buff.Bytes()
}

func (pow *ProofOfWork) PreparetoMine() []byte {
	info := [][]byte{
		pow.block.PrevBlockHash,
		pow.block.Data,
		//uint2byte(nonce),
		uint2byte(pow.block.Version),
		uint2byte(pow.block.Difficulty),
		[]byte(pow.block.TimeStamp),
		pow.block.MerkelRoot,
	}
	allinfo := bytes.Join(info, []byte{})
	hash := sha256.Sum256(allinfo)
	return hash[:]
}

//pow挖矿方法,返回两个参数，一个是碰撞成功的数字nonce，另一个是当前区块哈希值
func (pow *ProofOfWork) Mine() (uint64, []byte) {

	var nonce uint64
	var hash []byte
	hash = pow.PreparetoMine()
	/*
		for {
			hash = pow.PreparetoMine(nonce)
			var hashint big.Int
			hashint.SetBytes(hash)
			//less than
			if hashint.Cmp(pow.target) == -1 {
				break
			}
			nonce++
		}*/
	return nonce, hash
}
