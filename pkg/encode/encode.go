package encode

import (
	"fmt"

	"go-huffman/pkg/heap"
	"go-huffman/pkg/parse"
	"go-huffman/pkg/tree"
)

func Encode(msg string) {
	nodes := parse.Parse(msg)
	heap := heap.New()

	// make heap
	for _, n := range nodes {
		heap.Push(n)
	}

	// make tree
	for heap.Len() > 1 {
		l, r := heap.Pop(), heap.Pop()
		newNode := tree.NewInternal(l.Val() + r.Val())
		newNode.Left = l.(tree.Node)
		newNode.Right = r.(tree.Node)
		heap.Push(newNode)
	}

	// todo handle edge case of 0 or one nodes

	table := make(map[byte]string)
	traverse(heap.Pop().(tree.Node), "", &table)
	for b, s := range table {
		fmt.Printf("byte: %c, code: %s\n", b, s)
	}
}

func traverse(node tree.Node, str string, table *map[byte]string) {
	if node == nil {
		return
	}
	switch n := node.(type) {
	case *tree.InternalNode:
		traverse(n.Left, str+"0", table)
		traverse(n.Right, str+"1", table)
	case *tree.LeafNode:
		(*table)[n.Value] = str
	}
}
