package diffsquares

func SquareOfSum(n int) int {
	var res int
	for i := 0; i <= n; i++ {
		res += i
	}
	return res * res
}

func SumOfSquares(n int) int {
	var res int
	for i := 0; i <= n; i++ {
		res += i * i
	}
	return res
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
