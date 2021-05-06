// Package variablelengthquantity implements a solution to the problem on the Go track of Exercism
package variablelengthquantity

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/bits"
)

// EncodeVarint implements variable length quantity encoding.
func EncodeVarint(input []uint32) (res []byte) {
	for _, v := range input {
		//enc := encode(v)
		enc := encodeUint32(v)
		res = append(res, enc...)
	}
	return
}

// Using string representation of the bits
// https://stackoverflow.com/a/23192263/6454752
// https://stackoverflow.com/a/48647052/6454752
func encode(num uint32) []byte {
	// convert the bits of uint32 to the string representation
	bitString := fmt.Sprintf("%b", num)
	// find the needed length of the final byte array
	l := len(bitString)
	lenB := (l + l/7) / 8
	if lenB == 0 || (l+l/7)%8 != 0 {
		lenB++
	}
	bs := make([]byte, lenB)
	// iterate over the string in chunks of 7 backwards
	pos, cnt := l, 7
	var now byte
	for i := lenB - 1; i >= 0; i-- {
		if pos-cnt < 0 {
			cnt = pos
		}
		// convert each chunk to a byte
		now = strToByte(bitString[pos-cnt : pos])
		// set the leftmost bit on each chunk except the last one
		if i < lenB-1 {
			now = setBit(now, 7)
		}
		bs[i] = now
		pos -= 7
	}
	return bs
}

func strToByte(str string) (now byte) {
	for _, r := range str {
		now = now<<1 + byte(r-'0')
	}
	return
}

func setBit(n byte, pos uint) byte {
	n |= 1 << pos
	return n
}

func clearBit(n byte, pos uint) byte {
	return n &^ (1 << pos)
}

// Bit shifting implementation
func encodeUint32(num uint32) (res []byte) {
	// find the minimum number of bits needed to represent the input
	l := bits.Len32(num)
	// find the length of the byte array
	resLen := (l + l/7) / 8
	if resLen == 0 || (l+l/7)%8 != 0 {
		resLen++
	}

	if l < 28 {
		b := writeToBytes(num)
		res = append(res, clearBit(b[0], 7))
		for i := 1; i < resLen; i++ {
			num = num << 1
			b = writeToBytes(num)
			res = append([]byte{setBit(b[i], 7)}, res...)
		}
	} else {
		// cat to uint64 if the resulting number is larger than uint32
		num64 := uint64(num)
		b := writeToBytes(num64)
		res = append(res, clearBit(b[0], 7))
		for i := 1; i < resLen; i++ {
			num64 = num64 << 1
			b = writeToBytes(num64)
			res = append([]byte{setBit(b[i], 7)}, res...)
		}
	}
	return
}

func writeToBytes(num interface{}) []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, num)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	return buf.Bytes()
}

// DecodeVarint implements variable length quantity decoding.
func DecodeVarint(input []byte) (output []uint32, err error) {
	return
}
