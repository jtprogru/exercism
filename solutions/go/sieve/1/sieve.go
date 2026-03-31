package sieve

import "sort"

func Sieve(limit int) []int {
	A := map[int]bool{}
	var res []int

	if limit <= 1 {
		return res
	}

	for i := 2; i <= limit; i++ {
		A[i] = true
	}

	for i := 2; i <= limit*limit; i++ {
		if A[i] {
			for j := i * i; j <= limit; j += i {
				A[j] = false
			}
		}
	}
	for k, v := range A {
		if v {
			res = append(res, k)
		}
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})
	return res
}