package main

import (
	"go-huffman/pkg/encode"
)

func main() {
	err := encode.Encode("../../testfiles/small", "/tmp/out")
	if err != nil {
		panic(err)
	}
	err = encode.Decode("/tmp/out", "../../testfiles/out.txt")
	if err != nil {
		panic(err)
	}
}
