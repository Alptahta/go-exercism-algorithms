// Package raindrop implements number to rainbow converter
package raindrops

import (
	"strconv"
)

// Convert converts integer number into raindrop string
func Convert(n int) string {
	const n3 int = 3
	const n5 int = 5
	const n7 int = 7

	const hasFactor3 = "Pling"
	const hasFactor5 = "Plang"
	const hasFactor7 = "Plong"

	var raindrop string
	if n%n3 == 0 {
		raindrop = hasFactor3
	}
	if n%n5 == 0 {
		raindrop = raindrop + hasFactor5
	}
	if n%n7 == 0 {
		raindrop = raindrop + hasFactor7
	}
	if len(raindrop) == 0 {
		return strconv.Itoa(n)
	}
	return raindrop

}
