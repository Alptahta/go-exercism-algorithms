// Package diffsquares implements functions to calculate sum of squares
// squares of sum and difference between them
package diffsquares

// SquareOfSum calculates squares of sum of numbers to n
func SquareOfSum(n int) int {
	squareSum := (n * (n + 1)) / 2
	return squareSum * squareSum
}

// SumOfSquares calculates sum of squares of numbers to n
func SumOfSquares(n int) int {
	return (n * (n + 1) * (2*n + 1)) / 6
}

// Difference calculates difference between two functions at the above
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
