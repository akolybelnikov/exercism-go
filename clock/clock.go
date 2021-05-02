// Package clock implements a solution to the problem on Exercism
package clock

import "fmt"

const (
	minutes = 60
	hours   = 24
	day     = 1440
)

// Clock holds the time in minutes
type Clock int

// New constructs a new Clock
func New(h, m int) Clock {
	// Store time as a number of minutes.
	// Convert negative minutes to positive as difference to minutes in day.
	// If minutes are positive, we actually store minutes + day.
	return Clock((((h*minutes + m) % day) + day)%day)
}

// String converts the Clock to the readable time format
func (c Clock) String() string {
	h := (c/minutes)%hours
	m := c % minutes
	return fmt.Sprintf("%02d:%02d", h, m)
}

// Add returns the new state of the Clock after adding the minutes
func (c Clock) Add(m int) Clock {
	n := int(c) + m
	return New(n/minutes, n%minutes)
}

// Subtract returns the new state of the Clock after subtracting the minutes
func (c Clock) Subtract(m int) Clock {
	n := int(c) - m
	return New(n/minutes, n%minutes)
}
