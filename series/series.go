// Package series implements a solution to the problem on Go track of Exercism
package series

// All outputs all the contiguous substrings of length n in a string of digits.
func All(n int, s string) (result []string) {
	if n > len(s) {
		return
	}
	for i := 0; i <= len(s)-n; i++ {
		result = append(result, s[i:i+n])
	}
	return
}

// UnsafeFirst returns the first substring of s with length n.
func UnsafeFirst(n int, s string) string {
	return s[:n]
}

// First signals that in some cases you can't take the first N characters of the string.
func First(n int, s string) (first string, ok bool) {
	if n > len(s) {
		return
	}
	return s[:n], true
}
