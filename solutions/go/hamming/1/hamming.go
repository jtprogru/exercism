package hamming

import "errors"

func Distance(a, b string) (int, error) {
	var counter int

	if len(a) != len(b) {
		return counter, errors.New("strings of unequal length")
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			counter++
		}
	}
	return counter, nil
}
