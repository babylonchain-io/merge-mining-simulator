# Mockminer

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
$ ./mockminer
```