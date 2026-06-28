package resistorcolor

var colors = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

// Colors returns the list of all colors.
func Colors() []string {
	colorNames := make([]string, 0, len(colors))
	for k := range colors {
		colorNames = append(colorNames, k)
	}
	return colorNames
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	code, ok := colors[color]
	if !ok {
		return -1
	}
	return code
}
