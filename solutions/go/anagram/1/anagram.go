package anagram

import (
	"sort"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	subject = strings.ToLower(subject)
	var matches []string
	for _, candidate := range candidates {
		if isAnagram(subject, strings.ToLower(candidate)) {
			matches = append(matches, candidate)
		}
	}

	return matches
}

func isAnagram(sub, cand string) bool {
	return sub != cand && alphagram(sub) == alphagram(cand)
}

func alphagram(str string) string {
	chars := strings.Split(str, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}
