// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind string

const (
	// Pick values for the following identifiers used by the test program.
	NaT Kind = "not a triangle" // not a triangle
	Equ Kind = "equilateral"    // equilateral
	Iso Kind = "isosceles"      // isosceles
	Sca Kind = "scalene"        // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	var k Kind
	if a > b {
		a, b = b, a
	}
	if b > c {
		b, c = c, b
	}
	if a > b {
		a, b = b, a
	}
	switch {
	case a <= 0:
		k = NaT
	case a+b < c:
		k = NaT
	case a == c:
		k = Equ
	case a == b || b == c:
		k = Iso
	default:
		k = Sca
	}

	return k
}
