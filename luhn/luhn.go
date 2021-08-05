// Package luhn implements Valid function to fulfil luhn algorithm
package luhn

import (
	"log"
	"strconv"
	"strings"
)

// Valid validates given string with luhn algorithm
func Valid(s string) bool {
	s = strings.ReplaceAll(s, " ", "")
	_, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	if len(s) <= 1 {
		return len(s) > 1
	}

	var is []int
	for _, v := range s {
		// Check if rune is a number's unicode
		if v >= 48 && v <= 57 {
			// Convert rune to string and then convert it to integer
			i, err := strconv.Atoi(string(v))
			if err != nil {
				log.Fatal(err)
			}
			is = append(is, i)
		}
	}

	total := 0
	for i, v := range is {
		if i%2 == len(is)%2 {
			if v*2 > 9 {
				total += (v * 2) - 9
			} else {
				total += v * 2
			}
		} else {
			total += v
		}
	}

	if total%10 == 0 {
		return total%10 == 0
	}

	return false
}
