package tree

type Node interface{}

type Nodeheader struct {
	Freq  uint64
	Left  Node
	Right Node
}
