package etl

import (
	"strings"
)

func Transform(in map[int][]string) map[string]int {
	res := make(map[string]int)

	for idx, v := range in {
		for _, c := range v {
			key := strings.ToLower(c)
			res[key] = idx
		}
	}

	return res
}
