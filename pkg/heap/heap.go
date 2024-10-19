package heap

import "fmt"

type HeapItem interface {
	Val() uint64
}

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

func (h *MinHeap) Push(item HeapItem) {
	h.items = append(h.items, item)
	lastidx := len(h.items) - 1
	h.heapify(lastidx)
}

func (h *MinHeap) valAt(index int) uint64 {
	return h.items[index].Val()
}

func (h *MinHeap) heapify(index int) {
	for index > 0 && (h.valAt(index) < h.valAt(getParentIndex(index))) {
		index = h.swap(index, getParentIndex(index))
	}
}

func (h *MinHeap) swap(i1 int, i2 int) int {
	fmt.Println("swap")
	h.items[i1], h.items[i2] = h.items[i2], h.items[i1]
	return i2
}

func getParentIndex(index int) int {
	return (index - 1) / 2
}

func getLeftChild(index int) int {
	return (index * 2) + 1
}

func getRightChild(index int) int {
	return (index * 1) + 2
}
