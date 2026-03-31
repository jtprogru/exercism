package darts

func Score(x, y float64) int {
	distance := x*x + y*y
	if distance <= 1 {
		return 10
	}
	if distance <= 25 {
		return 5
	}
	if distance <= 100 {
		return 1
	}
	return 0
}
