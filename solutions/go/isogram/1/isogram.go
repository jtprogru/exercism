package isogram

import "strings"

func IsIsogram(word string) bool {
	lowerWord := strings.ToLower(word)
	for k, v := range lowerWord {
		if v != ' ' && v != '-' && k < strings.LastIndex(lowerWord, string(v)) {
			return false
		}
	}
	return true
}
