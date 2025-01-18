package main

import (
	"go-huffman/pkg/encode"
)

func main() {
	encode.Encode("../../testfiles/3mb.txt", "/tmp/out")
}
