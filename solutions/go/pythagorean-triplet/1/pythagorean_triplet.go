package pythagorean

type Triplet [3]int

func pyth(a, b, c int) bool {
	return a*a+b*b == c*c
}

// Range generates list of all Pythagorean triplets with side lengths
// in the provided range.
func Range(min, max int) []Triplet {
	var res []Triplet
	for a := min; a <= max; a++ {
		for b := a; b <= max; b++ {
			for c := b; c <= max; c++ {
				if pyth(a, b, c) {
					res = append(res, Triplet{a, b, c})
				}
			}
		}
	}

	return res
}

// Sum returns a list of all Pythagorean triplets with a certain perimeter.
func Sum(p int) []Triplet {
	var res []Triplet
	max := p / 2
	for a := 1; a <= max; a++ {
		for b := a; b <= max; b++ {
			if c := p - a - b; pyth(a, b, c) {
				res = append(res, Triplet{a, b, c})
			}
		}
	}
	return res
}
