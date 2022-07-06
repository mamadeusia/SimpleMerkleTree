package merkletree

import (
	"encoding/hex"

	"golang.org/x/crypto/sha3"
)

type Hashable interface {
	hash() Hash
}

type Hash [32]byte

func (h Hash) String() string {
	return hex.EncodeToString(h[:])
}
func hash(data []byte) Hash {
	hash := sha3.NewLegacyKeccak256()
	//hash.Write([]byte{0xcc})
	hash.Write(data)
	val := hash.Sum(nil)
	var output Hash
	copy(output[:], val[0:32])
	return output
}
