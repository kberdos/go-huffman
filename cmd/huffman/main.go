package main

import (
	"fmt"

	"go-huffman/pkg/heap"
	"go-huffman/pkg/parse"
)

func main() {
	nodes := parse.Parse("hello!")
	heap := heap.New()
	for _, n := range nodes {
		heap.Push(n)
	}
	fmt.Println(heap)
}
