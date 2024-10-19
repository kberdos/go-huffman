package tree

import "fmt"

type InternalNode struct {
	Nodeheader
}

func NewInternal(c uint64) *InternalNode {
	return &InternalNode{
		Nodeheader: Nodeheader{Freq: c},
	}
}

func (node *InternalNode) Val() uint64 {
	return node.Freq
}

func (node *InternalNode) String() string {
	return fmt.Sprintf("Count: %d", node.Freq)
}
