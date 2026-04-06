// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

// Hey should have a comment documenting it.
func Hey(greeting string) string {
	greeting = strings.TrimSpace(greeting)

	isYelling := anyCheck(greeting, unicode.IsUpper) && !anyCheck(greeting, unicode.IsLower)
	isQuestion := len(greeting) > 0 && greeting[len(greeting)-1] == '?'

	switch {
	case greeting == "":
		return "Fine. Be that way!"
	case isYelling && isQuestion:
		return "Calm down, I know what I'm doing!"
	case isYelling:
		return "Whoa, chill out!"
	case isQuestion:
		return "Sure."
	default:
		return "Whatever."
	}
}

/*anyCheck determines if anyCheck items in a string pass some test*/
func anyCheck(items string, test func(rune) bool) bool {
	for _, item := range items {
		if test(item) {
			return true
		}
	}
	return false
}
