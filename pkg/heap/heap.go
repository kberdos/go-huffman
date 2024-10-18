package heap

type HeapItem interface {
	Val() (val interface{}, err error)
}

type MinHeap struct {
	items []HeapItem
}
