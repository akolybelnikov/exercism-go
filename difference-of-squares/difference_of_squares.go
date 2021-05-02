// Package diffsquares implements a solution to the problem on Go track of the Exercism
package diffsquares

// SquareOfSum returns the square of the sum of the first n natural numbers
// The formula for finding the sum of first n numbers is n(n+1) / 2
func SquareOfSum(n int) int {
	s := n * (n + 1) / 2
	return s * s
}

// SumOfSquares returns the sum of the squares of the first n natural numbers
// The formula for finding the sum of squares of first n numbers is n(n+1)(2n+1) / 6
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// Difference returns the difference between the square of the sum and the sum of the squares of the first N natural numbers.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
