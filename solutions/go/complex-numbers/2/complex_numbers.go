package complexnumbers

import "math"

// Number represents a complex number re + im*i.
type Number struct {
	re float64
	im float64
}

func (n Number) Real() float64 {
	return n.re
}

func (n Number) Imaginary() float64 {
	return n.im
}

func (n Number) Add(n2 Number) Number {
	return Number{n.re + n2.re, n.im + n2.im}
}

func (n Number) Subtract(n2 Number) Number {
	return Number{n.re - n2.re, n.im - n2.im}
}

func (n Number) Multiply(n2 Number) Number {
	return Number{
		re: n.re*n2.re - n.im*n2.im,
		im: n.im*n2.re + n.re*n2.im,
	}
}

func (n Number) Times(factor float64) Number {
	return Number{n.re * factor, n.im * factor}
}

func (n Number) Divide(n2 Number) Number {
	denom := n2.re*n2.re + n2.im*n2.im
	return Number{
		re: (n.re*n2.re + n.im*n2.im) / denom,
		im: (n.im*n2.re - n.re*n2.im) / denom,
	}
}

func (n Number) Conjugate() Number {
	return Number{n.re, -n.im}
}

func (n Number) Abs() float64 {
	return math.Sqrt(n.re*n.re + n.im*n.im)
}

func (n Number) Exp() Number {
	ea := math.Exp(n.re)
	return Number{
		re: ea * math.Cos(n.im),
		im: ea * math.Sin(n.im),
	}
}
