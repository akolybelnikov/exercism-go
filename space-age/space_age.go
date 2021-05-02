// Package space provides solution to Exercism problem Space Age
package space

import "math"

type Planet string

// Age calculates and returns how old someone would be on a given planet, given an age in seconds
func Age(seconds float64, planet Planet) float64 {
	var earthSeconds = 31557600.0

	planets := make(map[Planet]float64)

	planets["Earth"] = 1.0
	planets["Mercury"] = 0.2408467
	planets["Venus"] = 0.61519726
	planets["Mars"] = 1.8808158
	planets["Jupiter"] = 11.862615
	planets["Saturn"] = 29.447498
	planets["Uranus"] = 84.016846
	planets["Neptune"] = 164.79132

	return math.Round((seconds/(earthSeconds*planets[planet]))*100) / 100
}
