// Package variablelengthquantity implements a solution to the problem on the Go track of Exercism
package variablelengthquantity

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// EncodeVarint implements variable length quantity encoding.
func EncodeVarint(input []uint32) []byte {
	buf := new(bytes.Buffer)
	for _, v := range input {
		err := binary.Write(buf, binary.LittleEndian, v)
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Printf("%032b\n", input)
	return buf.Bytes()
}

// DecodeVarint implements variable length quantity decoding.
func DecodeVarint(input []byte) (output []uint32, err error) {
	return
}
