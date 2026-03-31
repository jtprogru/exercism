package summultiples

func SumMultiples(limit int, divisors ...int) int {
	var res int
	for i := 1; i < limit; i++ {
		for _, d := range divisors {
			if d != 0 && i%d == 0 {
				res += i
				break
			}
		}
	}
	return res
}
