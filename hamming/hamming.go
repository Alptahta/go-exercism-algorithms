//Package hamming implements Hamming Distance
package hamming

import (
	"errors"
	"unicode/utf8"
)

// Distance calculates lenght of hamming distance of given two string
func Distance(a, b string) (int, error) {
	bsa := []rune(a)
	bsb := []rune(b)

	hammingDistance := 0

	if utf8.RuneCountInString(a) != utf8.RuneCountInString(b) {
		return hammingDistance, errors.New("length of strands should be equal")
	}
	for i, v := range bsa {
		if v != bsb[i] {
			hammingDistance++
		}
	}
	return hammingDistance, nil
}
