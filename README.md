# Merge-Mining-Simulator


Golang implementation of merge mining simulator.

## Overview
This project provides a merge-mining simulator, with a focus on the RPC interaction between a mining pool simulator and a bbl simulator. 

![avatar](mining.png)

The workfolw is shown as follows.

1.	A mining pool periodically collects bbld chains’ recent block header hashes and difficulty requirements.
2.	A mining pool computes an auxiliary Merkle root using these hashes and other data that the mining pool wants to inject, then puts this root into a certain field of the Bitcoin block it is going to mine, such as the sigScript field in the coinbase transaction. 

3.	once a worker in the mining pool mines a Bitcoin block whose difficulty matches a bbld chain’s requirement, return this block header and the inclusion proof of the bbld chain’s block header hash as the auxPoW.

4.	If a worker is mining more than one, it uses the following algorithm to convert the chain ID to a slot at the base of the merkle tree in which that chain's block hash must slot. Note that it is only a recommended solution, which has ready adopted by Namecoin and Elastos.

```
unsigned int rand = merkle_nonce;
rand = rand * 1103515245 + 12345;
rand += chain_id;
rand = rand * 1103515245 + 12345;
slot_num = rand % merkle_size
```

## Simulators

### Bbl node Simulator

```
$ cd ./sim-bbl-node
```

Sim-bbl-node is used to capture the functionalities of the Babylon node.

- [x] Bbl block hash preparation
- [x] key-value database
- [x] Createauxblock RPC service
- [x] Submitauxblock RPC service
- [x] Auxpow handle (e.g., auxpow verification)
- [ ] Mining confirmation 
- [ ] Peer to peer network

### BTC miner simulator

```
$ cd ./sim-merge-miner
```

Sim-merge-miner is used to capture the functionalities of the mining pool.

- [x] BTC nonce preparation
- [x] RPC periodically calling
- [ ] Mining confirmation 
- [ ] peer to peer network

## RPC Interfaces

### createauxblock

Generate an auxiliary block


| name         | type   | description     |
| ------------ | ------ | --------------- |
| paytoaddress | string | miner's address |


### submitauxblock

Submit the solved auxpow of an auxiliary block
 

| name      | type   | description                               |
| --------- | ------ | ----------------------------------------- |
| blockhash | string | the auxiliary block hash                  |
| auxpow    | string | the solved auxpow of this auxiliary block |
