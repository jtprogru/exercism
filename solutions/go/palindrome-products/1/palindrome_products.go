package palindromeproducts

import "errors"

// Product holds a palindrome value together with all factor pairs (a, b)
// within the requested range such that a*b equals the palindrome.
type Product struct {
	Product        int
	Factorizations [][2]int
}

// Products returns the smallest and largest palindrome products of two factors
// in the inclusive range [fmin, fmax], each with all of its factor pairs.
func Products(fmin, fmax int) (Product, Product, error) {
	if fmin > fmax {
		return Product{}, Product{}, errors.New("min must be <= max")
	}

	smallest := Product{}
	largest := Product{}
	found := false

	for a := fmin; a <= fmax; a++ {
		for b := a; b <= fmax; b++ {
			p := a * b
			if !isPalindrome(p) {
				continue
			}

			if !found {
				smallest = Product{p, [][2]int{{a, b}}}
				largest = Product{p, [][2]int{{a, b}}}
				found = true
				continue
			}

			switch {
			case p < smallest.Product:
				smallest = Product{p, [][2]int{{a, b}}}
			case p == smallest.Product:
				smallest.Factorizations = append(smallest.Factorizations, [2]int{a, b})
			}

			switch {
			case p > largest.Product:
				largest = Product{p, [][2]int{{a, b}}}
			case p == largest.Product:
				largest.Factorizations = append(largest.Factorizations, [2]int{a, b})
			}
		}
	}

	return smallest, largest, nil
}

// isPalindrome reports whether n reads the same forwards and backwards in base 10.
func isPalindrome(n int) bool {
	if n < 0 {
		return false
	}
	reversed, remaining := 0, n
	for remaining > 0 {
		reversed = reversed*10 + remaining%10
		remaining /= 10
	}
	return reversed == n
}
