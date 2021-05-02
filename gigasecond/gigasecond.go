// Package gigasecond provides a solution to the problem on Go track of Exercism.
package gigasecond

import "time"

// AddGigasecond determines the moment that would be after a gigasecond has passed given a moment as a starting point.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1000000000)
}
