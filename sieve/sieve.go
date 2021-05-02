// Package sieve implements a solution to the problem on Exercism
package sieve

import "math"

// Sieve uses the Sieve of Eratosthenes to find all the primes from 2 up to a given number.
func Sieve(n int) (result []int) {
	if n <= 1 {
		return []int{}
	}

	sieve := make([]bool, n+1)
	sqr := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqr; i++ {
		if sieve[i] == false {
			for j := i*i; j <= n; j += i {
				sieve[j] = true
			}
		}
	}
	for i := 2; i <=n; i++ {
		if !sieve[i] {
			result = append(result, i)
		}
	}

	return
}
