package pangram

import "strings"

func IsPangram(input string) bool {
	if len(input) == 0 {
		return false
	}
	var count int
	var check [26]bool
	for _, char := range strings.ToLower(input) {
		if char >= 'a' {
			if char > 'z' {
				continue
			}
			char -= 97
		} else {
			continue
		}
		if !check[char] {
			if count == 25 {
				return true
			}
			check[char] = true
			count++
		}
	}

	return false
}
