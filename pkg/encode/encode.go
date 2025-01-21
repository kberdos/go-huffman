package encode

import (
	"encoding/binary"
	"encoding/json"
	"os"

	"go-huffman/pkg/heap"
	"go-huffman/pkg/parse"
	"go-huffman/pkg/tree"

	bitio "github.com/kberdos/go-bitio"
)

// compresses the src file to the dst file
func Encode(src, dst string) error {
	file, err := os.Open(src)
	if err != nil {
		return err
	}
	fileinfo, err := file.Stat()
	if err != nil {
		return err
	}
	inbuffer := make([]byte, fileinfo.Size())
	_, err = file.Read(inbuffer)
	if err != nil {
		return err
	}

	nodes := parse.Parse(inbuffer)
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

	table := make(map[byte]encoding)
	traverse(heap.Pop().(tree.Node), encoding{0, 0}, table)

	outfile, err := os.Create(dst)
	if err != nil {
		return err
	}

	// tablesize (uint64) | table | encoding
	tablebytes, err := json.Marshal(table)
	if err != nil {
		return err
	}
	sizebuf := make([]byte, 8)
	binary.BigEndian.PutUint64(sizebuf, uint64(len(tablebytes)))
	tablebytes = append(sizebuf, tablebytes...)

	_, err = outfile.Write(tablebytes)
	if err != nil {
		return err
	}

	bw := bitio.NewWriter(outfile)

	for _, b := range inbuffer {
		bw.WriteBits(table[b].R, table[b].N)
	}
	err = bw.Close()
	if err != nil {
		return err
	}
	return nil
}

func traverse(node tree.Node, enc encoding, table map[byte]encoding) {
	if node == nil {
		return
	}
	switch n := node.(type) {
	case *tree.InternalNode:
		traverse(n.Left, enc.push0(), table)
		traverse(n.Right, enc.push1(), table)
	case *tree.LeafNode:
		table[n.Value] = enc
	}
}
