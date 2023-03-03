package merkletree

import (
	"bytes"
	"encoding/hex"
)

type Leaf string

type Node struct {
	left        Hashable
	right       Hashable
	inProofTree bool
}

type EmptyLeaf struct {
}

func (b Leaf) GetHash() Hash {
	data, _ := hex.DecodeString(string(b))
	return GetHash([]byte(data)[:])
}

func (_ EmptyLeaf) GetHash() Hash {
	return [32]byte{}
}

func (n Node) GetHash() Hash {
	var l, r [32]byte
	l = n.left.GetHash()
	r = n.right.GetHash()
	return GetHash(append(l[:], r[:]...))
}

func NewNode(left Hashable, right Hashable) Node {
	leftHash := left.GetHash()
	rightHash := right.GetHash()
	g := bytes.Compare(leftHash[:], rightHash[:])
	if g <= 0 {
		return Node{left: left, right: right}
	} else {
		return Node{left: right, right: left}
	}
}

type LeafSorted []Leaf

func (l LeafSorted) Len() int { return len(l) }

func (l LeafSorted) Less(i, j int) bool {
	left := l[i].GetHash()
	right := l[j].GetHash()
	g := bytes.Compare(left[:], right[:])
	if g <= 0 {
		return true
	} else {
		return false
	}
}

func (l LeafSorted) Swap(i, j int) { l[i], l[j] = l[j], l[i] }
