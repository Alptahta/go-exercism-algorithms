// Package raindrop implements number to rainbow converter
package raindrops

import "strconv"

type raindrop struct {
	divisor int
	drop    string
}

var drops = []raindrop{
	{3, "Pling"},
	{5, "Plang"},
	{7, "Plong"},
}

// Convert converts integer number into raindrop string
func Convert(n int) string {
	var raindrop string

	for _, drip := range drops {
		if n%drip.divisor == 0 {
			raindrop += drip.drop
		}
	}
	if raindrop == "" {
		raindrop = strconv.Itoa(n)
	}
	return raindrop

}
