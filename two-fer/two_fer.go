// Package twofer implements a simple library for Two-fer or 2-fer.
package twofer

import "fmt"

// ShareWith implements Two-fer and returns string.
func ShareWith(name string) string {
	if name == "" {
		return "One for you, one for me."
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
