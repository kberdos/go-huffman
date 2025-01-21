package encode

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/kberdos/go-bitio"
)

// decompress the src file to the dst file
func Decode(src, dst string) error {
	infile, err := os.Open(src)
	if err != nil {
		return err
	}
	sizebuf := make([]byte, 8)
	// read the size
	_, err = infile.Read(sizebuf)
	if err != nil {
		return err
	}
	tablesize := binary.BigEndian.Uint64(sizebuf)
	fmt.Printf("the table size is: %d\n", tablesize)
	tablebuf := make([]byte, tablesize)
	_, err = infile.Read(tablebuf)
	if err != nil {
		return err
	}
	bytetoenc := make(map[byte]encoding)
	err = json.Unmarshal(tablebuf, &bytetoenc)
	if err != nil {
		return err
	}
	fmt.Printf("size of decoded map is: %d\n", len(bytetoenc))
	enctobyte := reversetable(bytetoenc)
	fmt.Printf("size of reversed map is: %d\n", len(enctobyte))
	err = parsefromtable(infile, dst, enctobyte)
	if err != nil {
		return err
	}

	return nil
}

func reversetable(table map[byte]encoding) map[encoding]byte {
	out := make(map[encoding]byte)
	for k, v := range table {
		out[v] = k
	}
	return out
}

func parsefromtable(r io.Reader, dst string, table map[encoding]byte) error {
	outfile, err := os.Create(dst)
	if err != nil {
		return err
	}
	br, err := bitio.NewReader(r)
	if err != nil {
		return err
	}
	enc := encoding{0, 0}
	for {
		bit, err := br.ReadBits(1)
		if err != nil {
			fmt.Println("done")
			break
		}
		fmt.Println(bit)
		if bit == 1 {
			enc = enc.push1()
		} else {
			enc = enc.push0()
		}
		if b, ok := table[enc]; ok {
			fmt.Printf("found a match! %b\n", b)
			buf := make([]byte, 1)
			buf[0] = b
			n, err := outfile.Write(buf)
			if n != 1 {
				panic("what")
			}
			if err != nil {
				return err
			}
			enc = encoding{0, 0}
		}
	}
	err = outfile.Close()
	if err != nil {
		return err
	}

	return nil
}
