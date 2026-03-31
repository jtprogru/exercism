package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	freq := Frequency{}
	for _, word := range strings.Fields(normalize(phrase)) {
		word = strings.Trim(word, "'")
		freq[word]++
	}
	return freq
}

var reNormalize = regexp.MustCompile(`[^\w|']`)

func normalize(phrase string) string {
	phrase = strings.ToLower(phrase)
	return reNormalize.ReplaceAllLiteralString(phrase, " ")
}
