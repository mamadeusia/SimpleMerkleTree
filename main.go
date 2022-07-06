package main

import (
	"github.com/mamadeusia/keccak256MerkleTreeGo/merkletree"
)

func main() {
	// 3,4,5,7,10,11,12

	input := []string{"3fDfAb896ddC97E4CdFd44297F1d64172EEe761D", "17c114d0439Aed3f94c0E12F63Ae61e3BA70d4f1",
		"E7A817bbD2A4D30058b7fd9041ABA1Db3552cd8c", "733e2457cE010026eB9e658C501c851f0136E12a",
		"7A5c4CAd57b127454b3021371bB7f1F168c14c13", "B680e27775591AbB23622c4B33E96eF9d5A284F8"}
	root := merkletree.BuildTree(input)
	merkletree.PrintRoot(root)
	// merkletree.PrintTree(root)

	proofs := merkletree.GetProof("E7A817bbD2A4D30058b7fd9041ABA1Db3552cd8c", input)
	merkletree.PrintfProof(proofs)

}
