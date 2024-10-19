package main

import (
	"fmt"

	"go-huffman/pkg/heap"
	"go-huffman/pkg/parse"
)

func main() {
	nodes := parse.Parse("abbaacc !acb")
	heap := heap.New()
	for _, n := range nodes {
		heap.Push(n)
	}

	for item := heap.Pop(); item != nil; {
		fmt.Println(item)
		item = heap.Pop()
	}
}
