package bottlesong

import (
	"fmt"
)

var words = []string{"no", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}

func Recite(startBottles, takeDown int) []string {
	var result []string
	for i := range takeDown {
		if i > 0 {
			result = append(result, "")
		}
		result = append(result, verse(startBottles-i)...)
	}
	return result
}

func verse(n int) []string {
	current := Title(words[n])
	next := words[n-1]
	return []string{
		fmt.Sprintf("%s green %s hanging on the wall,", current, bottleWord(n)),
		fmt.Sprintf("%s green %s hanging on the wall,", current, bottleWord(n)),
		"And if one green bottle should accidentally fall,",
		fmt.Sprintf("There'll be %s green %s hanging on the wall.", next, bottleWord(n-1)),
	}
}

func bottleWord(n int) string {
	if n == 1 {
		return "bottle"
	}
	return "bottles"
}
