package sieve

func Sieve(limit int) []int {
	A := make([]bool, limit+1)
	res := []int{}
	for p := 2; p <= limit; {
		for i := p + p; i <= limit; i += p {
			A[i] = true
		}
		for p++; p <= limit && A[p]; p++ {
		}
	}
	for i := 2; i <= limit; i++ {
		if !A[i] {
			res = append(res, i)
		}
	}
	return res
}
