package merkletree

import (
	"fmt"
	"strings"
)

func PrintTree(node Node) {
	printNode(node, 0)
}

func PrintRoot(node Node) {
	fmt.Printf("0x%s \n", node.hash())
}

func PrintfProof(proofs []Hashable) {
	fmt.Printf("[")
	for i, proof := range proofs {
		fmt.Printf("\"0x%s\"", proof.hash())
		if i != len(proofs)-1 {
			fmt.Printf(",")
		} else {
			fmt.Printf("]\n")
		}
	}
}

func printNode(node Node, level int) {
	fmt.Printf("(%d) %s %s\n", level, strings.Repeat("-", level), node.hash())
	if l, ok := node.left.(Node); ok {
		printNode(l, level+1)
	} else if l, ok := node.left.(Leaf); ok {
		fmt.Printf("(%d) %s %s (data: %s)\n", level+1, strings.Repeat("-", level+1), l.hash(), l)
	}
	if r, ok := node.right.(Node); ok {
		printNode(r, level+1)
	} else if r, ok := node.right.(Leaf); ok {
		fmt.Printf("(%d) %s %s (data: %s)\n", level+1, strings.Repeat("-", level+1), r.hash(), r)
	}
}
