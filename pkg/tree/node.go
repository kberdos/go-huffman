package tree

import "go-huffman/pkg/heap"

type Node interface {
	heap.HeapItem
	String() string
}

type Nodeheader struct {
	Freq  uint64
	Left  Node
	Right Node
}
