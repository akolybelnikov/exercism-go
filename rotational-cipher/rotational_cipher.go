// Package rotationalcipher implements a solution to the problem on Go track of Exercism
package rotationalcipher

import (
	"bytes"
)

// RotationalCipher is an implementation of the rotational cipher, also sometimes called the Caesar cipher.
func RotationalCipher(s string, key int) string {
	if key == 0 || key == 26 {
		return s
	}
	var buf bytes.Buffer
	for _, r := range s {
		switch {
		// ascii codes for capital letters 65-90 and for small case letters 97-122
		case r > 64 && r < 91:
			r = int32((int(r)-65+key)%26 + 65)
			buf.WriteRune(r)
		case r > 96 && r < 123:
			r = int32((int(r)-97+key)%26 + 97)
			buf.WriteRune(r)
		default:
			buf.WriteRune(r)
		}
	}
	return buf.String()
}
