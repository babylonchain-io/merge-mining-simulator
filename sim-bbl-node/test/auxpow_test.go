package test

import (
	"bytes"
	"mockbbld/auxpow"
	"mockbbld/common"
	"testing"
)

// test auxpow deserialization, expected true
func TestDeserialize_Assert_True(t *testing.T) {
	var auxpow1 auxpow.AuxPow
	auxPowHex := "02000000010000000000000000000000000000000000000000000000000000000000000000ffffffff4b0313ee0904a880495b742f4254432e434f4d2ffabe6d6d9581ba0156314f1e92fd03430c6e4428a32bb3f1b9dc627102498e5cfbf26261020000004204cb9a010f32a00601000000000000ffffffff0200000000000000001976a914c0174e89bd93eacd1d5a1af4ba1802d412afc08688ac0000000000000000266a24aa21a9ede2f61c3f71d1defd3fa999dfa36953755c690689799962b48bebd836974e8cf90000000014acac4ee8fdd8ca7e0b587b35fce8c996c70aefdf24c333038bdba7af531266000000000001ccc205f0e1cb435f50cc2f63edd53186b414fcb22b719da8c59eab066cf30bdb0000000000000020d1061d1e456cae488c063838b64c4911ce256549afadfc6a4736643359141b01551e4d94f9e8b6b03eec92bb6de1e478a0e913e5f733f5884857a7c2b965f53ca880495bffff7f20a880495b"
	buf, _ := common.HexStringToBytes(auxPowHex)
	if err := auxpow1.Deserialize(bytes.NewReader(buf)); err != nil {
		t.Error("auxpow1 deserialization failed")
	}
}

// test auxpow deserialization, expected false
func TestDeserialize_Assert_False(t *testing.T) {
	var auxpow1 auxpow.AuxPow
	auxPowHex1 := "02000000000000000000000000000000000000000000000000000000000ffffffff4b0313ee0904a880495b742f4254432e434f4d2ffabe6d6d9581ba0156314f1e92fd03430c6e4428a32bb3f1b9dc627102498e5cfbf26261020000004204cb9a010f32a00601000000000000ffffffff0200000000000000001976a914c0174e89bd93eacd1d5a1af4ba1802d412afc08688ac0000000000000000266a24aa21a9ede2f61c3f71d1defd3fa999dfa36953755c690689799962b48bebd836974e8cf90000000014acac4ee8fdd8ca7e0b587b35fce8c996c70aefdf24c333038bdba7af531266000000000001ccc205f0e1cb435f50cc2f63edd53186b414fcb22b719da8c59eab066cf30bdb0000000000000020d1061d1e456cae488c063838b64c4911ce256549afadfc6a4736643359141b01551e4d94f9e8b6b03eec92bb6de1e478a0e913e5f733f5884857a7c2b965f53ca880495bffff7f20a880495b"
	buf1, _ := common.HexStringToBytes(auxPowHex1)
	if err := auxpow1.Deserialize(bytes.NewReader(buf1)); err != nil {
		t.Error("auxpow2 deserialization failed")
	}
}

// test difficulty, expected true
func TestAuxpowDifficulty_Assert_True(t *testing.T) {
	var auxpow auxpow.AuxPow
	normalAuxPowHex := "02000000010000000000000000000000000000000000000000000000000000000000000000ffffffff4b0313ee0904a880495b742f4254432e434f4d2ffabe6d6d9581ba0156314f1e92fd03430c6e4428a32bb3f1b9dc627102498e5cfbf26261020000004204cb9a010f32a00601000000000000ffffffff0200000000000000001976a914c0174e89bd93eacd1d5a1af4ba1802d412afc08688ac0000000000000000266a24aa21a9ede2f61c3f71d1defd3fa999dfa36953755c690689799962b48bebd836974e8cf90000000014acac4ee8fdd8ca7e0b587b35fce8c996c70aefdf24c333038bdba7af531266000000000001ccc205f0e1cb435f50cc2f63edd53186b414fcb22b719da8c59eab066cf30bdb0000000000000020d1061d1e456cae488c063838b64c4911ce256549afadfc6a4736643359141b01551e4d94f9e8b6b03eec92bb6de1e478a0e913e5f733f5884857a7c2b965f53ca880495bffff7f20a880495b"
	buf, _ := common.HexStringToBytes(normalAuxPowHex)
	auxpow.Deserialize(bytes.NewReader(buf))

	// check the Difficulty
	targetDifficulty := common.CompactToBig(auxpow.ParBlockHeader.Bits)
	hash := auxpow.ParBlockHeader.Hash()
	// hash should be less that targetDifficulty
	if common.HashToBig(&hash).Cmp(targetDifficulty) > 0 {
		t.Error("auxpow checking failed, difficulty is not satified for bbld's requirement")
	}
}

// test difficulty, expected false
func TestAuxpowDifficulty_Assert_False(t *testing.T) {
	var auxpow auxpow.AuxPow
	normalAuxPowHex := "02000000010000000000000000000000000000000000000000000000000000000000000000ffffffff4b0313ee0904a880495b742f4254432e434f4d2ffabe6d6d9581ba0156314f1e92fd03430c6e4428a32bb3f1b9dc627102498e5cfbf26261020000004204cb9a010f32a00601000000000000ffffffff0200000000000000001976a914c0174e89bd93eacd1d5a1af4ba1802d412afc08688ac0000000000000000266a24aa21a9ede2f61c3f71d1defd3fa999dfa36953755c690689799962b48bebd836974e8cf90000000014acac4ee8fdd8ca7e0b587b35fce8c996c70aefdf24c333038bdba7af531266000000000001ccc205f0e1cb435f50cc2f63edd53186b414fcb22b719da8c59eab066cf30bdb0000000000000020d1061d1e456cae488c063838b64c4911ce256549afadfc6a4736643359141b01551e4d94f9e8b6b03eec92bb6de1e478a0e913e5f733f5884857a7c2b965f53ca880495bffff7f20a880495b"
	buf, _ := common.HexStringToBytes(normalAuxPowHex)
	auxpow.Deserialize(bytes.NewReader(buf))

	targetDifficulty := common.CompactToBig(auxpow.ParBlockHeader.Bits)
	// here, we set a random nocne.
	auxpow.ParBlockHeader.Nonce = uint32(0000)
	hash := auxpow.ParBlockHeader.Hash()
	// hash should be less that targetDifficulty
	if common.HashToBig(&hash).Cmp(targetDifficulty) > 0 {
		t.Error("auxpow checking failed, difficulty is not satified for bbld's requirement")
	}
}

// test MerkleRoot in BtcHeader
func TestAuxBlockHashInMerkleRoot_Assert_True(t *testing.T) {
	var auxpow auxpow.AuxPow
	normalAuxPowHex := "02000000010000000000000000000000000000000000000000000000000000000000000000ffffffff4b0313ee0904a880495b742f4254432e434f4d2ffabe6d6d9581ba0156314f1e92fd03430c6e4428a32bb3f1b9dc627102498e5cfbf26261020000004204cb9a010f32a00601000000000000ffffffff0200000000000000001976a914c0174e89bd93eacd1d5a1af4ba1802d412afc08688ac0000000000000000266a24aa21a9ede2f61c3f71d1defd3fa999dfa36953755c690689799962b48bebd836974e8cf90000000014acac4ee8fdd8ca7e0b587b35fce8c996c70aefdf24c333038bdba7af531266000000000001ccc205f0e1cb435f50cc2f63edd53186b414fcb22b719da8c59eab066cf30bdb0000000000000020d1061d1e456cae488c063838b64c4911ce256549afadfc6a4736643359141b01551e4d94f9e8b6b03eec92bb6de1e478a0e913e5f733f5884857a7c2b965f53ca880495bffff7f20a880495b"
	buf, _ := common.HexStringToBytes(normalAuxPowHex)
	auxpow.Deserialize(bytes.NewReader(buf))

	// check if coinbase is in btc block header
	if auxpow.AuxBlockHashInMerkleRoot() != true {
		t.Error("auxpow checking failed, merkle root failed, coinbase is not in btc block header")
	}
}

// test MerkleRoot in BtcHeader
func TestAuxBlockHashInMerkleRoot_Assert_False(t *testing.T) {
	var auxpow auxpow.AuxPow
	normalAuxPowHex := "02000000010000000000000000000000000000000000000000000000000000000000000000ffffffff4b0313ee0904a880495b742f4254432e434f4d2ffabe6d6d9581ba0156314f1e92fd03430c6e4428a32bb3f1b9dc627102498e5cfbf26261020000004204cb9a010f32a00601000000000000ffffffff0200000000000000001976a914c0174e89bd93eacd1d5a1af4ba1802d412afc08688ac0000000000000000266a24aa21a9ede2f61c3f71d1defd3fa999dfa36953755c690689799962b48bebd836974e8cf90000000014acac4ee8fdd8ca7e0b587b35fce8c996c70aefdf24c333038bdba7af531266000000000001ccc205f0e1cb435f50cc2f63edd53186b414fcb22b719da8c59eab066cf30bdb0000000000000020d1061d1e456cae488c063838b64c4911ce256549afadfc6a4736643359141b01551e4d94f9e8b6b03eec92bb6de1e478a0e913e5f733f5884857a7c2b965f53ca880495bffff7f20a880495b"
	buf, _ := common.HexStringToBytes(normalAuxPowHex)
	auxpow.Deserialize(bytes.NewReader(buf))

	//here, we set a CoinBaseMerkle.
	auxpow.ParCoinBaseMerkle = make([]common.Uint256, 0)
	if auxpow.AuxBlockHashInMerkleRoot() != true {
		t.Error("auxpow checking failed, merkle root failed, coinbase is not in btc block header")
	}
}

// test MerkleRootInCoinbase
func TestMerkleRootInCoinbase_Assert_True(t *testing.T) {

	var auxpow auxpow.AuxPow
	normalAuxPowHex := "02000000010000000000000000000000000000000000000000000000000000000000000000ffffffff4b0313ee0904a880495b742f4254432e434f4d2ffabe6d6d9581ba0156314f1e92fd03430c6e4428a32bb3f1b9dc627102498e5cfbf26261020000004204cb9a010f32a00601000000000000ffffffff0200000000000000001976a914c0174e89bd93eacd1d5a1af4ba1802d412afc08688ac0000000000000000266a24aa21a9ede2f61c3f71d1defd3fa999dfa36953755c690689799962b48bebd836974e8cf90000000014acac4ee8fdd8ca7e0b587b35fce8c996c70aefdf24c333038bdba7af531266000000000001ccc205f0e1cb435f50cc2f63edd53186b414fcb22b719da8c59eab066cf30bdb0000000000000020d1061d1e456cae488c063838b64c4911ce256549afadfc6a4736643359141b01551e4d94f9e8b6b03eec92bb6de1e478a0e913e5f733f5884857a7c2b965f53ca880495bffff7f20a880495b"
	buf, _ := common.HexStringToBytes(normalAuxPowHex)
	auxpow.Deserialize(bytes.NewReader(buf))

	var hashAuxBlock *common.Uint256
	blockHashHex := "7926398947f332fe534b15c628ff0cd9dc6f7d3ea59c74801dc758ac65428e64"
	hashAuxBlock, err := common.Uint256FromHexString(blockHashHex)
	if err != nil {
		t.Error("auxpow checking failed, hex string to uint256 failed")
	}

	// check if block is in Coinbase
	//hashAuxBlockBytes := common.BytesReverse(hashAuxBlock.Bytes())
	//hashAuxBlock, _ = common.Uint256FromBytes(hashAuxBlockBytes)

	auxRootHash := auxpow.GetMerkleRoot(*hashAuxBlock, auxpow.AuxMerkleBranch, auxpow.AuxMerkleIndex)

	if auxpow.MerkleRootInCoinbase(hashAuxBlock, &auxRootHash) != true {
		t.Error("auxpow checking failed,  merkle root is not in coinbase")
	}
}

// test MerkleRootInCoinbase
func TestMerkleRootInCoinbase_Assert_False(t *testing.T) {

	var auxpow auxpow.AuxPow
	normalAuxPowHex := "02000000010000000000000000000000000000000000000000000000000000000000000000ffffffff4b0313ee0904a880495b742f4254432e434f4d2ffabe6d6d9581ba0156314f1e92fd03430c6e4428a32bb3f1b9dc627102498e5cfbf26261020000004204cb9a010f32a00601000000000000ffffffff0200000000000000001976a914c0174e89bd93eacd1d5a1af4ba1802d412afc08688ac0000000000000000266a24aa21a9ede2f61c3f71d1defd3fa999dfa36953755c690689799962b48bebd836974e8cf90000000014acac4ee8fdd8ca7e0b587b35fce8c996c70aefdf24c333038bdba7af531266000000000001ccc205f0e1cb435f50cc2f63edd53186b414fcb22b719da8c59eab066cf30bdb0000000000000020d1061d1e456cae488c063838b64c4911ce256549afadfc6a4736643359141b01551e4d94f9e8b6b03eec92bb6de1e478a0e913e5f733f5884857a7c2b965f53ca880495bffff7f20a880495b"
	buf, _ := common.HexStringToBytes(normalAuxPowHex)
	auxpow.Deserialize(bytes.NewReader(buf))

	var hashAuxBlock *common.Uint256
	blockHashHex := "7926398947f332fe534b15c628ff0cd9dc6f7d3ea59c74801dc758ac65428e64"
	hashAuxBlock, err := common.Uint256FromHexString(blockHashHex)
	if err != nil {
		t.Error("auxpow checking failed, hex string to uint256 failed")
	}

	// check if block is in Coinbase
	hashAuxBlockBytes := common.BytesReverse(hashAuxBlock.Bytes())
	hashAuxBlock, _ = common.Uint256FromBytes(hashAuxBlockBytes)
	auxRootHash := common.Uint256{}

	if auxpow.MerkleRootInCoinbase(hashAuxBlock, &auxRootHash) != true {
		t.Error("auxpow checking failed,  merkle root is not in coinbase")
	}
}

// test AuxWorkInMiningMerkleTree
func TestAuxWorkInMiningMerkleTree_Assert_True(t *testing.T) {

	var auxpow auxpow.AuxPow
	normalAuxPowHex := "02000000010000000000000000000000000000000000000000000000000000000000000000ffffffff4b0313ee0904a880495b742f4254432e434f4d2ffabe6d6d9581ba0156314f1e92fd03430c6e4428a32bb3f1b9dc627102498e5cfbf26261020000004204cb9a010f32a00601000000000000ffffffff0200000000000000001976a914c0174e89bd93eacd1d5a1af4ba1802d412afc08688ac0000000000000000266a24aa21a9ede2f61c3f71d1defd3fa999dfa36953755c690689799962b48bebd836974e8cf90000000014acac4ee8fdd8ca7e0b587b35fce8c996c70aefdf24c333038bdba7af531266000000000001ccc205f0e1cb435f50cc2f63edd53186b414fcb22b719da8c59eab066cf30bdb0000000000000020d1061d1e456cae488c063838b64c4911ce256549afadfc6a4736643359141b01551e4d94f9e8b6b03eec92bb6de1e478a0e913e5f733f5884857a7c2b965f53ca880495bffff7f20a880495b"

	buf, _ := common.HexStringToBytes(normalAuxPowHex)
	auxpow.Deserialize(bytes.NewReader(buf))

	var hashAuxBlock *common.Uint256
	blockHashHex := "7926398947f332fe534b15c628ff0cd9dc6f7d3ea59c74801dc758ac65428e64"
	hashAuxBlock, err := common.Uint256FromHexString(blockHashHex)
	if err != nil {
		t.Error("auxpow checking failed, hex string to uint256 failed")
	}
	// check if auxwork is in mining merkle tree
	if !auxpow.AuxWorkInMiningMerkleTree(hashAuxBlock, 6, 1) {
		t.Error("auxpow checking failed, auxwork not in merkle tree")
	}
}

// test AuxWorkInMiningMerkleTree
func TestAuxWorkInMiningMerkleTree_Assert_False(t *testing.T) {

	var auxpow auxpow.AuxPow
	normalAuxPowHex := "02000000010000000000000000000000000000000000000000000000000000000000000000ffffffff4b0313ee0904a880495b742f4254432e434f4d2ffabe6d6d9581ba0156314f1e92fd03430c6e4428a32bb3f1b9dc627102498e5cfbf26261020000004204cb9a010f32a00601000000000000ffffffff0200000000000000001976a914c0174e89bd93eacd1d5a1af4ba1802d412afc08688ac0000000000000000266a24aa21a9ede2f61c3f71d1defd3fa999dfa36953755c690689799962b48bebd836974e8cf90000000014acac4ee8fdd8ca7e0b587b35fce8c996c70aefdf24c333038bdba7af531266000000000001ccc205f0e1cb435f50cc2f63edd53186b414fcb22b719da8c59eab066cf30bdb0000000000000020d1061d1e456cae488c063838b64c4911ce256549afadfc6a4736643359141b01551e4d94f9e8b6b03eec92bb6de1e478a0e913e5f733f5884857a7c2b965f53ca880495bffff7f20a880495b"

	buf, _ := common.HexStringToBytes(normalAuxPowHex)
	auxpow.Deserialize(bytes.NewReader(buf))

	var hashAuxBlock *common.Uint256
	blockHashHex := "7926398947f332fe534b15c628ff0cd9dc6f7d3ea59c74801dc758ac65428e64"
	hashAuxBlock, err := common.Uint256FromHexString(blockHashHex)
	if err != nil {
		t.Error("auxpow checking failed, hex string to uint256 failed")
	}
	// change the merkle height
	if !auxpow.AuxWorkInMiningMerkleTree(hashAuxBlock, 6, 2) {
		t.Error("auxpow checking failed, auxwork not in merkle tree")
	}
}
