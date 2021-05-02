package cryptosquare

import (
	"math"
	"unicode"
)

func Encode(pt string) string {
	if pt == "" {
		return ""
	}
	runes := normalize(pt)
	// find the number of columns and the length or rows
	c := int(math.Ceil(math.Sqrt(float64(len(runes)))))
	r := c
	if c*(r-1) >= len(runes) {
		r--
	}
	// pad the runes with whitespaces if their length is not divisible by the row length
	for c*r > len(runes) {
		runes = append(runes, 32)
	}
	var result []rune
	for i := 0; i < c; i++ {
		for j := 0; j < r; j++ {
			result = append(result, runes[j*c+i])
		}
		if i < c-1 {
			result = append(result, 32)
		}
	}

	return string(result)
}

func normalize(input string) (output []rune) {
	for _, r := range input {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			output = append(output, unicode.ToLower(r))
		}
	}
	return
}
