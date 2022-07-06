package merkletree

func GetProof(l string, leafs []string) []Hashable {
	var j int
	for i, val := range leafs {
		if l == val {
			j = i
		}

	}
	var output []Hashable
	var nodes []Node
	for i := 0; i < len(leafs); i += 2 {
		if i+1 < len(leafs) {
			node := NewNode(Leaf(leafs[i]), Leaf(leafs[i+1]))
			if j == i {
				output = append(output, Leaf(leafs[i+1]))
				node.inProofTree = true
			} else if j == i+1 {
				output = append(output, Leaf(leafs[i]))
				node.inProofTree = true
			}
			nodes = append(nodes, node)
		} else {
			node := NewNode(Leaf(leafs[i]), Leaf(leafs[i]))
			if j == i {
				output = append(output, Leaf(leafs[i]))
				node.inProofTree = true
			}
			nodes = append(nodes, node)
		}
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
