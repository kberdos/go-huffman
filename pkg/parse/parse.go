package parse

import (
	"go-huffman/pkg/tree"
)

// Parse a msg into leaf nodes
func Parse(msg []byte) []tree.Node {
	bytemap := make(map[byte]uint64)
	for _, b := range msg {
		bytemap[b] += 1
	}
	out := make([]tree.Node, 0)
	for b, c := range bytemap {
		out = append(out, tree.NewLeaf(b, c))
	}
	return out
}
