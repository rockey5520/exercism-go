package diffsquares

// SquareOfSum caluclated Squares of sum
func SquareOfSum(n int) int {

	return (n * (n + 1) / 2) * (n * (n + 1) / 2)
}

// SumOfSquares caluclated sum of squares
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// Difference caluclates difference between Squares of sum and sum of squares
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
