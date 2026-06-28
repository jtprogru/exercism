package bottlesong

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

var words = []string{"no", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}

// capitalize returns s with its first rune upper-cased.
func capitalize(s string) string {
	if s == "" {
		return s
	}
	r, size := utf8.DecodeRuneInString(s)
	return string(unicode.ToUpper(r)) + s[size:]
}

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
	current := capitalize(words[n])
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
