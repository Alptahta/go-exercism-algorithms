// Package grains implements grains for chessboard story
package grains

import "errors"

// Square calculates grains at specific square
func Square(n int) (uint64, error) {
	if n == 0 {
		return 0, errors.New("square 0 returns an error")
	}
	if n < 0 {
		return 0, errors.New("negative square returns an error")
	}
	if n > 64 {
		return 0, errors.New("square greater than 64 returns an error")
	}

	return IntPow(n - 1), nil

}

// Total calculates total number of grains in chessboard
func Total() uint64 {
	return IntPow(65) - 1
}

// IntPow calculates power of 2
func IntPow(m int) uint64 {
	if m == 0 {
		return 1
	}
	result := 2
	for i := 2; i <= m; i++ {
		result *= 2
	}
	return uint64(result)
}
