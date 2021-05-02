// Package reverse provides a solution to the problem on Go track of Exercism
package reverse

// Reverse accepts a string argument and returns a reversed version of it
func Reverse(input string) (result string) {
	result = ""
	var runes = make([]rune, 0, len(input))
	for _, r := range input {
		runes = append([]rune{r}, runes...)
	}
	if len(runes) > 0 {
		result = string(runes)
	}

	return
}
