//Package hamming implements Hamming Distance
package hamming

import "errors"

// Distance calculates lenght of hamming distance of given two string
func Distance(a, b string) (int, error) {
	bsa := []byte(a)
	bsb := []byte(b)

	hammingDistance := 0

	if len(a) != len(b) {
		return hammingDistance, errors.New("Length of strands should be equal")
	}
	for i, v := range bsa {
		if v == bsb[i] {
			hammingDistance += 0
		} else {
			hammingDistance++
		}
	}
	return hammingDistance, nil
}
