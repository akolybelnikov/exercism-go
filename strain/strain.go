// Package strain implements a solution to the problem on Go track of Exercism
package strain

// Ints is a collection of integers
type Ints []int

// Strings is a collection of strings
type Strings []string

// Lists is a collection of collections of integers
type Lists [][]int

// Keep returns new Ints containing those elements where the predicate func returns true
func (i Ints) Keep(f func(int) bool) (result Ints) {
	for _, el := range i {
		if ok := f(el); ok {
			result = append(result, el)
		}
	}
	return
}

// Discard returns new Ints containing those elements where the predicate func returns false
func (i Ints) Discard(f func(int) bool) (result Ints) {
	for _, el := range i {
		if ok := f(el); !ok {
			result = append(result, el)
		}
	}
	return
}

// Keep returns new Lists containing those elements where the predicate func returns true
func (l Lists) Keep(f func([]int) bool) (result Lists) {
	for _, el := range l {
		if ok := f(el); ok {
			result = append(result, el)
		}
	}
	return
}

// Keep returns new Strings containing those elements where the predicate func returns true
func (s Strings) Keep(f func(string) bool) (result Strings) {
	for _, el := range s {
		if ok := f(el); ok {
			result = append(result, el)
		}
	}
	return
}
