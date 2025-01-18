package encode

import (
	"go-huffman/pkg/heap"
	"go-huffman/pkg/parse"
	"go-huffman/pkg/tree"
)

func Encode(msg string) string { // for now does strings
	nodes := parse.Parse(msg)
	heap := heap.New()

	// make heap
	for _, n := range nodes {
		heap.Push(n)
	}

	// make tree
	for heap.Len() > 1 {
		l, r := heap.Pop(), heap.Pop() // lower freq on the left
		newNode := tree.NewInternal(l.Val() + r.Val())
		newNode.Left = l.(tree.Node)
		newNode.Right = r.(tree.Node)
		heap.Push(newNode)
	}

	// todo handle edge case of 0 or one nodes

	table := make(map[byte]string)
	traverse(heap.Pop().(tree.Node), "", table)

	encoded := ""
	for _, b := range msg {
		encoded += table[byte(b)]
	}
	return encoded
}

func traverse(node tree.Node, str string, table map[byte]string) {
	if node == nil {
		return
	}
	switch n := node.(type) {
	case *tree.InternalNode:
		traverse(n.Left, str+"0", table)
		traverse(n.Right, str+"1", table)
	case *tree.LeafNode:
		table[n.Value] = str
	}
}
