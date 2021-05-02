// Package summultiples implements a solution ot the problem on Go track of Exercism
package summultiples

// SumMultiples returns the sum of all the unique multiples of divisors up to but not including the limit
func SumMultiples(limit int, divisors ...int) (s int) {
	for i := 1; i < limit; i++ {
		for _, j := range divisors {
			if j != 0 && i%j == 0 {
				s += i
				break
			}
		}
	}
	return
}
