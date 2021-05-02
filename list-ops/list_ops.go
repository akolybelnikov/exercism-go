// Package listops provides a solution to the problem on Go track of exercism
package listops

// IntList is an integer slice type that implements the list ops
type IntList []int

// binFunc is a binary function used in ops with two integers, returns an integer
type binFunc func(x, y int) int

// predFunc accepts and integer argument and returns a boolean value
type predFunc func(n int) bool

// unaryFunc is used in ops with one integer, returns an integer
type unaryFunc func(x int) int

// Foldr reduces each item into the accumulator from the right using the binFunc
func (l IntList) Foldr(fn binFunc, acc int) int {
	for _, v := range l.Reverse() {
		acc = fn(v, acc)
	}
	return acc
}

// Foldl reduces each item into the accumulator from the left using the binFunc
func (l IntList) Foldl(fn binFunc, acc int) int {
	for _, v := range l {
		acc = fn(acc, v)
	}
	return acc
}

// Filter returns the list of all items for which predicate(item) is True
func (l IntList) Filter(fn predFunc) IntList {
	result := make(IntList, 0, l.Length())
	i := 0
	for _, v := range l {
		if fn(v) {
			result = result[:i+1]
			result[i] = v
			i++
		}
	}
	return result
}

// Length returns the total number of items within the list
func (l IntList) Length() (length int) {
	for range l {
		length++
	}
	return
}

// Map returns the list of the results of applying function(item) on all items
func (l IntList) Map(fn unaryFunc) IntList {
	result := make(IntList, l.Length())
	i := 0
	for _, v := range l {
		result[i] = fn(v)
		i++
	}
	return result
}

// Reverse returns a list with all the original items, but in reversed order
func (l IntList) Reverse() IntList {
	len := l.Length()
	result := make(IntList, len)
	for i := len - 1; i >= 0; i-- {
		result[len-1-i] = l[i]
	}
	return result
}

// Append adds all items in the second list to the end of the first list
func (l IntList) Append(s IntList) IntList {
	result := make(IntList, l.Length()+s.Length())
	i := 0
	for _, v := range l {
		result[i] = v
		i++
	}
	for _, v := range s {
		result[i] = v
		i++
	}

	return result
}

// Concat combines all items in all given lists into one flattened list
func (l IntList) Concat(args []IntList) IntList {
	for _, a := range args {
		l = l.Append(a)
	}
	return l
}
