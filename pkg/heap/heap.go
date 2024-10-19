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

func (h *MinHeap) Len() int {
	return len(h.items)
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
	h.insert_heapify(lastidx)
}

func (h *MinHeap) Pop() HeapItem {
	if len(h.items) == 0 {
		return nil
	}
	item := h.items[0]
	lastidx := len(h.items) - 1
	h.items[0] = h.items[lastidx]
	h.items = h.items[:lastidx]
	if len(h.items) > 0 {
		h.delete_heapify(0)
	}
	return item
}

func (h *MinHeap) valAt(index int) uint64 {
	return h.items[index].Val()
}

func (h *MinHeap) insert_heapify(index int) {
	for index > 0 && (h.valAt(index) < h.valAt(getParentIndex(index))) {
		index = h.swap(index, getParentIndex(index))
	}
}

func (h *MinHeap) get(index int) HeapItem {
	if index >= len(h.items) {
		return nil
	}
	return h.items[index]
}

func (h *MinHeap) delete_heapify(index int) {
	for {
		minIndex, minVal := index, h.items[index].Val()
		li, ri := getLeftChildIndex(index), getRightChildIndex(index)
		l, r := h.get(li), h.get(ri)
		if l != nil && l.Val() < minVal {
			minIndex, minVal = li, l.Val()
		}
		if r != nil && r.Val() < minVal {
			minIndex, minVal = ri, r.Val()
		}
		if minIndex != index {
			h.swap(index, minIndex)
			index = minIndex
		} else {
			break
		}
	}
}

func (h *MinHeap) swap(i1 int, i2 int) int {
	h.items[i1], h.items[i2] = h.items[i2], h.items[i1]
	return i2
}

func getParentIndex(index int) int {
	return (index - 1) / 2
}

func getLeftChildIndex(index int) int {
	return (index * 2) + 1
}

func getRightChildIndex(index int) int {
	return (index * 1) + 2
}
