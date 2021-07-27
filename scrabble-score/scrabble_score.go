// Package scrabble implements score function to calculate scrabble score
package scrabble

import "strings"

// I found two solutions. But could not decide which one to pick. Both of them need improvements i believe.

// FIRST SOLUTION

var characters = map[rune]int{
	'a': 1, 'e': 1, 'i': 1, 'o': 1, 'u': 1, 'l': 1, 'n': 1, 'r': 1, 's': 1, 't': 1,
	'd': 2, 'g': 2,
	'b': 3, 'c': 3, 'm': 3, 'p': 3,
	'f': 4, 'h': 4, 'v': 4, 'w': 4, 'y': 4,
	'k': 5,
	'j': 8, 'x': 8,
	'q': 10, 'z': 10,
}

// Score calculates scrabble score
func Score(s string) int {
	score := 0

	for _, r := range strings.ToLower(s) {
		score += characters[r]
	}

	return score

}

// SECOND SOLUTION

// type letterAndValues struct {
// 	letter []string
// 	value  int
// }

// var characters = []letterAndValues{
// 	{[]string{"a", "e", "i", "o", "u", "l", "n", "r", "s", "t"}, 1},
// 	{[]string{"d", "g"}, 2},
// 	{[]string{"b", "c", "m", "p"}, 3},
// 	{[]string{"f", "h", "v", "w", "y"}, 4},
// 	{[]string{"k"}, 5},
// 	{[]string{"j", "x"}, 8},
// 	{[]string{"q", "z"}, 10},
// }

// Score calculates scrabble score
// func Score(s string) int {
// 	ls := strings.ToLower(s)
// 	rs := []rune(ls)

// 	score := 0

// 	if s == "" {
// 		return score
// 	}
// 	for _, v1 := range rs {
// 		for _, v2 := range characters {
// 			for _, v3 := range v2.letter {
// 				if string(v1) == v3 {
// 					score += v2.value
// 				}
// 			}
// 		}
// 	}
// 	return score

// }
