// Package accumulate provides a solution to the problem on Go track of Exercism
package accumulate

// Accumulate returns a collection resulting from applying the given operation to each element of the input collection.
func Accumulate(given []string, converter func(string) string) (result []string) {
	result = []string{}
	if len(given) > 0 {
		for _, elem := range given {
			result = append(result, converter(elem))
		}
	}
	return
}
