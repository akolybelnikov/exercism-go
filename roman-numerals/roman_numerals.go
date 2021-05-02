// Package romannumerals implements a solution to the problem on Go track of Exercism
package romannumerals

import (
	"bytes"
	"errors"
	"strconv"
)

var romanNumerals = map[int]string{1: "I", 5: "V", 10: "X", 50: "L", 100: "C", 500: "D", 1000: "M"}

// ToRomanNumeral converts from normal numbers to Roman Numerals
func ToRomanNumeral(arabic int) (string, error) {
	if arabic > 0 && arabic <= 3000 {
		var b bytes.Buffer
		switch length(arabic) {
		case 4:
			b.WriteString(findRomanNumerals(arabic, 1000))
		case 3:
			b.WriteString(findRomanNumerals(arabic, 100))
		case 2:
			b.WriteString(findRomanNumerals(arabic, 10))
		case 1:
			b.WriteString(findRomanNumerals(arabic, 1))
		}
		return b.String(), nil
	}
	return "", errors.New("invalid input")
}

func length(n int) int {
	return len(strconv.Itoa(n))
}

func findRomanNumerals(n int, factor int) string {
	var b bytes.Buffer
	switch t := n / factor; {
	case t < 4:
		for i := 0; i < t; i++ {
			b.WriteString(romanNumerals[factor])
		}
	case t == 4:
		b.WriteString(romanNumerals[factor])
		b.WriteString(romanNumerals[factor*5])
	case t < 9:
		b.WriteString(romanNumerals[factor*5])
		for i := 0; i < t%5; i++ {
			b.WriteString(romanNumerals[factor])
		}
	case t == 9:
		b.WriteString(romanNumerals[factor])
		b.WriteString(romanNumerals[factor*10])
	}
	t := n % factor
	if t > 0 {
		if s, err := ToRomanNumeral(t); err == nil {
			b.WriteString(s)
		}
	}
	return b.String()
}
