package heap

import "fmt"

type HeapItem interface{}

type MinHeap struct {
	items []HeapItem
}

func New() *MinHeap {
	return &MinHeap{}
}

func (h *MinHeap) String() string {
	out := "Heap Items:\n"
	for _, item := range h.items {
		out += fmt.Sprintln(item)
	}
	return out
}

func (h *MinHeap) Push(item HeapItem) error {
	h.items = append(h.items, item)
	return nil
}
