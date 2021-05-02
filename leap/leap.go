// Package leap provides an implementation of the Leap year exercise on Exercism's Go track.
package leap

// IsLeapYear checks if a given year is leap and returns a boolean value.
func IsLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				return true
			}
			return false
		}
		return true
	}
	return false
}
