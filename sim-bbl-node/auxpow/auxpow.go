package auxpow

import (
	"encoding/binary"
	"encoding/hex"
	"io"
	"strings"

	"mockbbld/common"
	"mockbbld/logger"
)

var (
	AuxPowChainID         = 1224
	pchMergedMiningHeader = []byte{0xfa, 0xbe, 'm', 'm'}
)

type AuxPow struct {
	AuxMerkleBranch   []common.Uint256
	AuxMerkleIndex    int
	ParCoinbaseTx     BtcTx            //bitcoin transaction
	ParCoinBaseMerkle []common.Uint256 //The merkle branch linking the coinbase_txn to the parent block's merkle_root
	ParMerkleIndex    int
	ParBlockHeader    BtcHeader      //Parent block header
	ParentHash        common.Uint256 //
}

func NewAuxPow(AuxMerkleBranch []common.Uint256, AuxMerkleIndex int,
	ParCoinbaseTx BtcTx, ParCoinBaseMerkle []common.Uint256,
	ParMerkleIndex int, ParBlockHeader BtcHeader) *AuxPow {

	return &AuxPow{
		AuxMerkleBranch:   AuxMerkleBranch,
		AuxMerkleIndex:    AuxMerkleIndex,
		ParCoinbaseTx:     ParCoinbaseTx,
		ParCoinBaseMerkle: ParCoinBaseMerkle,
		ParMerkleIndex:    ParMerkleIndex,
		ParBlockHeader:    ParBlockHeader,
	}
}

func (ap *AuxPow) Serialize(w io.Writer) error {
	err := ap.ParCoinbaseTx.Serialize(w)
	if err != nil {
		return err
	}

	err = ap.ParentHash.Serialize(w)
	if err != nil {
		return err
	}

	count := uint64(len(ap.ParCoinBaseMerkle))
	err = common.WriteVarUint(w, count)
	if err != nil {
		return err
	}

	for _, pcbm := range ap.ParCoinBaseMerkle {
		err = pcbm.Serialize(w)
		if err != nil {
			return err
		}
	}

	index := uint32(ap.ParMerkleIndex)
	err = common.WriteUint32(w, index)
	if err != nil {
		return err
	}

	count = uint64(len(ap.AuxMerkleBranch))
	err = common.WriteVarUint(w, count)
	if err != nil {
		return err
	}

	for _, amb := range ap.AuxMerkleBranch {
		err = amb.Serialize(w)
		if err != nil {
			return err
		}
	}

	index = uint32(ap.AuxMerkleIndex)
	err = common.WriteUint32(w, index)
	if err != nil {
		return err
	}

	err = ap.ParBlockHeader.Serialize(w)
	if err != nil {
		return err
	}

	return nil
}

func (ap *AuxPow) Deserialize(r io.Reader) error {
	err := ap.ParCoinbaseTx.Deserialize(r)
	if err != nil {
		return err
	}

	err = ap.ParentHash.Deserialize(r)
	if err != nil {
		return err
	}

	count, err := common.ReadVarUint(r, 0)
	if err != nil {
		return err
	}

	ap.ParCoinBaseMerkle = make([]common.Uint256, 0)
	for i := uint64(0); i < count; i++ {
		var hash common.Uint256
		err = hash.Deserialize(r)
		if err != nil {
			return err
		}
		ap.ParCoinBaseMerkle = append(ap.ParCoinBaseMerkle, hash)
	}

	index, err := common.ReadUint32(r)
	if err != nil {
		return err
	}
	ap.ParMerkleIndex = int(index)

	count, err = common.ReadVarUint(r, 0)
	if err != nil {
		return err
	}

	ap.AuxMerkleBranch = make([]common.Uint256, 0)
	for i := uint64(0); i < count; i++ {
		var hash common.Uint256
		err = hash.Deserialize(r)
		if err != nil {
			return err
		}
		ap.AuxMerkleBranch = append(ap.AuxMerkleBranch, hash)
	}

	index, err = common.ReadUint32(r)
	if err != nil {
		return err
	}
	ap.AuxMerkleIndex = int(index)

	err = ap.ParBlockHeader.Deserialize(r)
	if err != nil {
		return err
	}

	return nil
}

func (ap *AuxPow) Check(blockHashHex string, chainID int) bool {

	var hashAuxBlock *common.Uint256

	hashAuxBlock, err := common.Uint256FromHexString(blockHashHex)
	if err != nil {
		logger.Error.Println("hex string to uint256 failed")
		return false
	}

	// check the Difficulty
	targetDifficulty := common.CompactToBig(ap.ParBlockHeader.Bits)
	hash := ap.ParBlockHeader.Hash()

	// hash should be less that targetDifficulty
	if common.HashToBig(&hash).Cmp(targetDifficulty) > 0 {
		logger.Error.Println("difficulty is not satified for bbld's requirement")
		return false
	}

	// check if coinbase is in btc block header
	if !ap.CoinbaseInBtcHeader() {
		logger.Error.Println("merkle root failed, coinbase is not in btc block header")
		return false
	}

	// check if block is in Coinbase
	if !ap.BlockHashInCoinbase(hashAuxBlock) {
		logger.Error.Println("hashAuxBlock is not in coinbase")
		return false
	}

	if !ap.AuxWorkInMerkleTree(hashAuxBlock, chainID) {
		logger.Error.Println("auxwork not in merkle tree")
		return false
	}

	return true
}

// Coinbase in Btc Header
func (ap *AuxPow) CoinbaseInBtcHeader() bool {
	return GetMerkleRoot(ap.ParCoinbaseTx.Hash(), ap.ParCoinBaseMerkle, ap.ParMerkleIndex) == ap.ParBlockHeader.MerkleRoot
}

// AuxBlock Hash is not in coinbase
func (ap *AuxPow) BlockHashInCoinbase(hashAuxBlock *common.Uint256) bool {

	//reverse the hashAuxBlock
	hashAuxBlockBytes := common.BytesReverse(hashAuxBlock.Bytes())
	hashAuxBlock, _ = common.Uint256FromBytes(hashAuxBlockBytes)
	auxRootHash := GetMerkleRoot(*hashAuxBlock, ap.AuxMerkleBranch, ap.AuxMerkleIndex)

	script := ap.ParCoinbaseTx.TxIn[0].SignatureScript
	scriptStr := hex.EncodeToString(script)

	// reverse the auxRootHash
	auxRootHashReverseStr := hex.EncodeToString(common.BytesReverse(auxRootHash.Bytes()))
	pchMergedMiningHeaderStr := hex.EncodeToString(pchMergedMiningHeader)

	headerIndex := strings.Index(scriptStr, pchMergedMiningHeaderStr)
	rootHashIndex := strings.Index(scriptStr, auxRootHashReverseStr)

	if (headerIndex == -1) || (rootHashIndex == -1) {
		logger.Error.Println("hashAuxBlock 260 is not in coinbase")
		return false
	}

	if strings.Index(scriptStr[headerIndex+2:], pchMergedMiningHeaderStr) != -1 {
		logger.Error.Println("hashAuxBlock 265 is not in coinbase")
		return false
	}

	if headerIndex+len(pchMergedMiningHeaderStr) != rootHashIndex {
		logger.Error.Println("hashAuxBlock 270 is not in coinbase")
		return false
	}

	rootHashIndex += len(auxRootHashReverseStr)
	if len(scriptStr)-rootHashIndex < 8 {
		logger.Error.Println("hashAuxBlock 276 is not in coinbase")
		return false
	}
	return true
}

// Aux work in the merkle tree
func (ap *AuxPow) AuxWorkInMerkleTree(hashAuxBlock *common.Uint256, chainID int) bool {

	hashAuxBlockBytes := common.BytesReverse(hashAuxBlock.Bytes())
	hashAuxBlock, _ = common.Uint256FromBytes(hashAuxBlockBytes)
	auxRootHash := GetMerkleRoot(*hashAuxBlock, ap.AuxMerkleBranch, ap.AuxMerkleIndex)

	script := ap.ParCoinbaseTx.TxIn[0].SignatureScript
	scriptStr := hex.EncodeToString(script)

	// reverse the auxRootHash
	auxRootHashReverseStr := hex.EncodeToString(common.BytesReverse(auxRootHash.Bytes()))
	rootHashIndex := strings.Index(scriptStr, auxRootHashReverseStr)

	//Aux work merkle tree
	size := binary.LittleEndian.Uint32(script[rootHashIndex/2 : rootHashIndex/2+4])
	merkleHeight := len(ap.AuxMerkleBranch)
	if size != uint32(1<<uint32(merkleHeight)) {
		//return false
	}

	nonce := binary.LittleEndian.Uint32(script[rootHashIndex/2+4 : rootHashIndex/2+8])
	if ap.AuxMerkleIndex != GetExpectedIndex(nonce, chainID, merkleHeight) {
		logger.Error.Println("expected index failed")
		return false
	}

	return true
}

func GetMerkleRoot(hash common.Uint256, merkleBranch []common.Uint256, index int) common.Uint256 {
	if index == -1 {
		return common.Uint256{}
	}
	var sha [64]byte
	for _, it := range merkleBranch {
		if (index & 1) == 1 {
			copy(sha[:32], it[:])
			copy(sha[32:], hash[:])
			hash = common.Hash(sha[:])
		} else {
			copy(sha[:32], hash[:])
			copy(sha[32:], it[:])
			hash = common.Hash(sha[:])
		}
		index >>= 1
	}
	return hash
}

func GetExpectedIndex(nonce uint32, chainID, h int) int {
	rand := nonce
	rand = rand*1103515245 + 12345
	rand += uint32(chainID)
	rand = rand*1103515245 + 12345
	return int(rand % (1 << uint32(h)))
}
