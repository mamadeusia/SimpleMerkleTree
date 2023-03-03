package merkletree

import (
	"fmt"
	"strings"
)

func PrintTree(node Node) {
	if node.right == node.left { //for the case that size of addresses==1
		PrintRoot(node)
	} else {
		printNode(node, 0)
	}
}

func PrintRoot(node Node) {
	if node.right == node.left { //for the case that size of addresses==1
		fmt.Printf("0x%s \n", node.right.GetHash())
	} else {
		fmt.Printf("0x%s \n", node.GetHash())
	}
}

func PrintfProof(proofs []Hashable) {
	fmt.Printf("[")
	for i, proof := range proofs {
		fmt.Printf("\"0x%s\"", proof.GetHash())
		if i != len(proofs)-1 {
			fmt.Printf(",")
		}
	}
	fmt.Printf("]\n")

}

func printNode(node Node, level int) {
	fmt.Printf("(%d) %s %s\n", level, strings.Repeat("-", level), node.GetHash())
	if l, ok := node.left.(Node); ok {
		printNode(l, level+1)
	} else if l, ok := node.left.(Leaf); ok {
		fmt.Printf("(%d) %s %s (data: %s)\n", level+1, strings.Repeat("-", level+1), l.GetHash(), l)
	}
	if r, ok := node.right.(Node); ok {
		printNode(r, level+1)
	} else if r, ok := node.right.(Leaf); ok {
		fmt.Printf("(%d) %s %s (data: %s)\n", level+1, strings.Repeat("-", level+1), r.GetHash(), r)
	}
}
