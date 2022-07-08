
# Simple Merkle Tree for Solidity and Golang

This repository contains On-chain merkle tree creator. You can add address to onchain and update the merkleRoot of the contract by calling `addToList(address newAddress)` that is only accessible by owner or admin of the contract.

Also you can find indexed variables in the log as I emitted 
`event NewRoot(address indexed NewAddress,bytes32 indexed NewRoot)`
in each call to `addToList` function . 

Offchain side is responsible for `GetProof` as main functionality, also you can `PrintRoot` and `PrintTree` for visualization. 
 

## OffChain usage

Add you're addresses that you want to have in the tree 
in `main.go` . 


```bash
    go run main.go
```
    
## Onchain usage 
You can use remixd to test the functionality of solidity code in remix . 


## Hints 
There are some optimization that can be done on contract side 
like using struct and use of smaller uint like `uint32` to optimize the usage of storage. 
## Authors

- [@mamadeusia](https://github.com/mamadeusia)
