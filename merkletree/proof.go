package merkletree

import "strings"

func GetProof(l string, leafs []string) []Hashable {
	var leafsTrimmed []string
	for _, leaf := range leafs {
		leafsTrimmed = append(leafsTrimmed, strings.TrimPrefix(leaf, "0x"))
	}

	var j int
	if len(leafsTrimmed) == 1 { //for the case on one leaf
		return []Hashable{}
	}
	for i, val := range leafsTrimmed {
		if l == val {
			j = i
		}
	}
	var output []Hashable
	var nodes []Node
	for i := 0; i < len(leafsTrimmed); i += 2 {
		if i+1 < len(leafsTrimmed) {
			node := NewNode(Leaf(leafsTrimmed[i]), Leaf(leafsTrimmed[i+1]))
			if j == i {
				output = append(output, Leaf(leafsTrimmed[i+1]))
				node.inProofTree = true
			} else if j == i+1 {
				output = append(output, Leaf(leafsTrimmed[i]))
				node.inProofTree = true
			}
			nodes = append(nodes, node)
		} else {
			node := NewNode(Leaf(leafsTrimmed[i]), Leaf(leafsTrimmed[i]))
			if j == i {
				output = append(output, Leaf(leafsTrimmed[i]))
				node.inProofTree = true
			}
			nodes = append(nodes, node)
		}
	}
	if len(nodes) == 1 {
		return []Hashable{output[0]}
	}
	return appendProof(nodes, output)
}

func appendProof(parts []Node, input []Hashable) []Hashable {

	output := input
	var newParts []Node
	for i := 0; i < len(parts); i += 2 {
		if i+1 < len(parts) {
			node := NewNode(parts[i], parts[i+1])
			if parts[i].inProofTree {
				output = append(output, parts[i+1])
				node.inProofTree = true
			} else if parts[i+1].inProofTree {
				output = append(output, parts[i])
				node.inProofTree = true
			}
			newParts = append(newParts, node)
		} else {
			node := NewNode(parts[i], parts[i])
			if parts[i].inProofTree {
				output = append(output, parts[i])
				node.inProofTree = true
			}
			newParts = append(newParts, node)
		}
	}
	if len(newParts) == 1 {
		return output
	} else if len(newParts) > 1 {
		return appendProof(newParts, output)
	} else {
		panic("huh?!")
	}
}
