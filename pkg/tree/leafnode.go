package tree

import "fmt"

type LeafNode struct {
	Nodeheader
	Value byte
}

func NewLeaf(b byte, c uint64) *LeafNode {
	return &LeafNode{
		Value:      b,
		Nodeheader: Nodeheader{Freq: c},
	}
}

func (node *LeafNode) String() string {
	return fmt.Sprintf("Char: %c, Count: %d", node.Value, node.Freq)
}
