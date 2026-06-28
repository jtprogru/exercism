package sieve

func Sieve(limit int) []int {
	A := make([]bool, limit+1)
	res := []int{}
	for p := 2; p <= limit; p++ {
		if A[p] {
			continue
		}
		res = append(res, p)
		for i := p + p; i <= limit; i += p {
			A[i] = true
		}
	}
	return res
}
