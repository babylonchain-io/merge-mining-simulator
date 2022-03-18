# Mockbbld

## Requirements
Go 1.16 or newer.

Ensure Go was installed properly and is a supported version:

```
$ go version
$ go env GOROOT GOPATH
```

NOTE: The GOROOT and GOPATH above must not be the same path. It is recommended that GOPATH is set to a directory in your home directory such as `~/goprojects` to avoid write permission issues. It is also recommended to add `$GOPATH/bin` to your PATH at this point.

## Installation

### Download the code:

```
$ go clone https://github.com/babylonchain-io/mockbbld
$ cp -rf mockbbld/* $GOPATH/src/github.com/
```

### Build
```
$ go build
```

### Run
```
$ ./mockbbld

root@ip-172-31-21-176:~/go/src/github.com/mockbbld# ./mockbbld 
2022/03/17 08:51:54 Starting RPC Server on :10000
Version : 0
PrevBlockHash : 4f38694fe801802bc39567102159acb22244746e57d0c0cd161846b2f4635f19
Hash : 513a330875eed46d66f26912f573d5b4a9dfc4648b493c44235bdb07b5cc8c1f
MerkleRoot : 
TimeStamp : 2022-08:51:59
Difficuty : 5
Nonce : 0

```

### Test
```
curl -X POST -H "Content-Type: application/json" -d "{\"method\":\"BBLService.Createauxblock\",\"params\":[{\"paytoaddress\":\"Ef4UcaHwvFrFzzsyVf5YH4JBWgYgUqfTAB\"}],\"id\":\"1\"}" http://127.0.0.1:10000/rpc
```

```
curl -X POST -H "Content-Type: application/json" -d "{\"method\":\"BBLService.SubmitAuxBlock\",\"params\":[{\"blockhash\":\"fed7452b57d3bb6faf4122a54fa7fb6cb8a7f2727a10dcf7af5cc49d247bb22e\", \"auxpow\": \"02000000010000000000000000000000000000000000000000000000000000000000000000ffffffff4b0313ee0904a880495b742f4254432e434f4d2ffabe6d6d9581ba0156314f1e92fd03430c6e4428a32bb3f1b9dc627102498e5cfbf26261020000004204cb9a010f32a00601000000000000ffffffff0200000000000000001976a914c0174e89bd93eacd1d5a1af4ba1802d412afc08688ac0000000000000000266a24aa21a9ede2f61c3f71d1defd3fa999dfa36953755c690689799962b48bebd836974e8cf90000000014acac4ee8fdd8ca7e0b587b35fce8c996c70aefdf24c333038bdba7af531266000000000001ccc205f0e1cb435f50cc2f63edd53186b414fcb22b719da8c59eab066cf30bdb0000000000000020d1061d1e456cae488c063838b64c4911ce256549afadfc6a4736643359141b01551e4d94f9e8b6b03eec92bb6de1e478a0e913e5f733f5884857a7c2b965f53ca880495bffff7f20a880495b\"}],\"id\":\"1\"}" http://127.0.0.1:10000/rpc
```

Reference: <https://github.com/elastos/Elastos.ELA/>