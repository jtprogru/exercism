// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"slices"
	"strings"
	"unicode"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	var acronym strings.Builder
	for _, word := range SplitAny(CleanString(s), []string{" ", "-"}) {
		acronym.WriteString(strings.ToUpper(string(word[0])))
	}
	return acronym.String()
}

func SplitAny(s string, seps []string) []string {
	var result []string
	word := ""
	for _, char := range s {
		isSep := slices.Contains(seps, string(char))
		if isSep {
			if word != "" {
				result = append(result, word)
				word = ""
			}
		} else {
			word += string(char)
		}
	}
	if word != "" {
		result = append(result, word)
	}
	return result
}

func CleanString(input string) string {
	var builder strings.Builder
	builder.Grow(len(input))

	for _, r := range input {
		if unicode.IsLetter(r) || unicode.IsSpace(r) || r == '-' {
			builder.WriteRune(r)
		}
	}

	return builder.String()
}
