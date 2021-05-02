// Package prime implements a solution to the problem on the Go track of Exercism
package prime

// Nth determines what the nth prime is
func Nth(n int) (int, bool) {
	if n == 0 {
		return 0, false
	}
	// using incremental sieve: primes = [2, 3, ...] \ [[p², p²+p, ...] for p in primes],
	// interleave generation of primes with their squares.
	primes := []int{2}
	squares := make([]int, n)
	squares[0] = 4
	i := 0
	// maintain a separate index for squares
	j := 0
	for len(primes) < n {
		prime := false
		next := primes[i] + 1
		for !prime {
			if next == squares[j] {
				j++
			}
			if isPrime(next, primes[:j+1]) {
				i++
				primes = append(primes, next)
				squares[i] = next * next
				prime = true
			} else {
				next++
			}
		}
	}
	return primes[len(primes)-1], true
}

// primes are generated through divisibility testing by sequential primes
func isPrime(n int, primes []int) bool {
	for _, p := range primes {
		if n%p == 0 {
			return false
		}
	}
	return true
}
