package merkletree

import "strings"

func BuildTree(addresses []string) Node {
	var addressesTrimmed []string
	for _, address := range addresses {
		addressesTrimmed = append(addressesTrimmed, strings.TrimPrefix(address, "0x"))
	}

	var nodes []Hashable
	for i := 0; i < len(addresses); i += 2 {
		if i+1 < len(addresses) {
			nodes = append(nodes, NewNode(Leaf(addressesTrimmed[i]), Leaf(addressesTrimmed[i+1])))
		} else {
			nodes = append(nodes, NewNode(Leaf(addressesTrimmed[i]), Leaf(addressesTrimmed[i])))
		}
	}
	var output []Hashable
	if len(nodes) == 1 {
		return nodes[0].(Node)
	} else {
		output = buildTree(nodes)
	}
	return output[0].(Node)
}

func buildTree(parts []Hashable) []Hashable {

	var nodes []Hashable
	var i int
	for i = 0; i < len(parts); i += 2 {
		if i+1 < len(parts) {
			nodes = append(nodes, NewNode(parts[i], parts[i+1]))
		} else {
			nodes = append(nodes, NewNode(parts[i], parts[i]))
		}
	}
	if len(nodes) == 1 {
		return nodes
	} else if len(nodes) > 1 {
		return buildTree(nodes)
	} else {
		panic("huh?!")
	}
}
