package rotationalcipher

import (
	"unicode"
)

func RotationalCipher(plain string, shiftKey int) string {
	var res []rune
	for _, char := range plain {
		if unicode.IsUpper(char) {
			char = rune('A' + (int(char-'A')+shiftKey)%26)
		} else if unicode.IsLower(char) {
			char = rune('a' + (int(char-'a')+shiftKey)%26)
		}
		res = append(res, char)
	}

	return string(res)
}
