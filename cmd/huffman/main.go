package main

import (
	"go-huffman/pkg/encode"
)

func main() {
	encode.Encode("../../testfiles/small", "/tmp/out")
}
