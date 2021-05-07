// Package variablelengthquantity implements a solution to the problem on the Go track of Exercism
package variablelengthquantity

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"math/bits"
)

const (
	BITS = 8
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
	lenB := (l + l/(BITS-1)) / BITS
	if lenB == 0 || (l+l/(BITS-1))%BITS != 0 {
		lenB++
	}
	bs := make([]byte, lenB)
	// iterate over the string in chunks of 7 backwards
	pos, cnt := l, BITS-1
	var now byte
	for i := lenB - 1; i >= 0; i-- {
		if pos-cnt < 0 {
			cnt = pos
		}
		// convert each chunk to a byte
		now = strToByte(bitString[pos-cnt : pos])
		// set the leftmost bit on each chunk except the last one
		if i < lenB-1 {
			now = setBit(now, BITS-1)
		}
		bs[i] = now
		pos -= BITS - 1
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
	resLen := (l + l/(BITS-1)) / BITS
	if resLen == 0 || (l+l/(BITS-1))%BITS != 0 {
		resLen++
	}

	if l < 28 {
		b := writeToBytes(num)
		res = append(res, clearBit(b[0], BITS-1))
		for i := 1; i < resLen; i++ {
			num = num << 1
			b = writeToBytes(num)
			res = append([]byte{setBit(b[i], BITS-1)}, res...)
		}
	} else {
		// cat to uint64 if the resulting number is larger than uint32
		num64 := uint64(num)
		b := writeToBytes(num64)
		res = append(res, clearBit(b[0], BITS-1))
		for i := 1; i < resLen; i++ {
			num64 = num64 << 1
			b = writeToBytes(num64)
			res = append([]byte{setBit(b[i], BITS-1)}, res...)
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
	count := 0
	var words [][]byte
	var word []byte
	for i := 0; i < len(input); i++ {
		if count == 0 {
			word = []byte{}
		}
		word = append(word, input[i])
		if bits.LeadingZeros8(input[i]) == 0 {
			count++
		} else {
			words = append(words, word)
			count = 0
		}
	}

	if count != 0 {
		return []uint32{}, errors.New("incomplete sequence")
	}

	for _, w := range words {
		output = append(output, decode(w))
	}

	return
}

func decode(bs []byte) uint32 {
	l := len(bs)
	var self byte
	res := make([]byte, 4)
	curr := 0
	for i := l - 1; i >= 0; i-- {
		self = bs[i]
		self = clearBit(self, BITS-1)
		pos := BITS - (l - 1 - i)
		self = self >> (BITS - pos)
		if i > 0 {
			for j := uint(pos); j <= BITS-1; j++ {
				self = clearBit(self, j)
			}
		}
		if i > 0 {
			self += bs[i-1] << (pos - 1)
		}
		if curr < 4 {
			res[curr] = self
			curr++
		}
	}
	return binary.LittleEndian.Uint32(res)
}
