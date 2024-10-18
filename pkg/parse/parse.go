package parse

import (
	"go-huffman/pkg/tree"
)

// Parse a msg into leaf nodes
func Parse(msg string) []tree.Node {
	bmap := make(map[byte]uint64)
	for _, b := range []byte(msg) {
		bmap[b] += 1
	}
	out := make([]tree.Node, 0)
	for b, c := range bmap {
		out = append(out, tree.NewLeaf(b, c))
	}
	return out
}
