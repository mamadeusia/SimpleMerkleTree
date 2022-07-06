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

func (b Leaf) hash() Hash {
	data, _ := hex.DecodeString(string(b))
	return hash([]byte(data)[:])
}

func (_ EmptyLeaf) hash() Hash {
	return [32]byte{}
}

func (n Node) hash() Hash {
	var l, r [32]byte
	l = n.left.hash()
	r = n.right.hash()
	return hash(append(l[:], r[:]...))
}

func NewNode(left Hashable, right Hashable) Node {
	leftHash := left.hash()
	rightHash := right.hash()
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
	left := l[i].hash()
	right := l[j].hash()
	g := bytes.Compare(left[:], right[:])
	if g <= 0 {
		return true
	} else {
		return false
	}
}

func (l LeafSorted) Swap(i, j int) { l[i], l[j] = l[j], l[i] }
