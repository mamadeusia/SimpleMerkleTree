
# Simple Merkle Tree for Solidity and Golang

This repository consists of simple golang merkle tree and related solidity code that 
you can work with it . 



## Usage

Add you're addresses that you want to have in the tree 
in `main.go` . after that you can call the `PrintProof` and `PrintRoot` to 
use in solidity contract . you can also see the structure of tree by 
`PrintTree` method . 

```bash
    go run main.go
```
    
## Demo

Find the deployed smart contract on BSC_Testnet in the address below.
`0x97c34750E7FDfE7777d953B7B3f3305599E293a2`.
this is the private key for the address `0xE7A817bbD2A4D30058b7fd9041ABA1Db3552cd8c`
`95889a9207b16aa6735c9c5825c70214e775fcf33978f51ade7f47984a877aae`
you can use it to check the functionality . 
addresses that used in creating current root are in `main.go`.
## Authors

- [@mamadeusia](https://github.com/mamadeusia)
