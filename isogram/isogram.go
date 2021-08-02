// Package isogram implements function for isogram check of given string
package isogram

import "strings"

// IsIsogram check if given string is isogram
func IsIsogram(s string) bool {
	rs := []rune(strings.ToLower(s))

	for _, v := range rs {
		// Check if v value is between Uncapitalized and Capitalized letter's unicodes
		if (v >= 65 && v <= 90) || (v >= 97 && v <= 122) {
			if containsDuplicate(rs, v) {
				return false
			}
		} else {
			continue
		}
	}
	return true
}

// containsDuplicate check if given rune exist more than once in given rune slice
func containsDuplicate(rs []rune, r rune) bool {
	count := 0
	for _, v := range rs {
		if r == v {
			count++
		}
	}
	if count > 1 {
		return count > 1
	}
	return false
}
