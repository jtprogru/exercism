package encode

import (
	"fmt"
	"strconv"
)

func RunLengthEncode(input string) string {
	count := 1
	var res string
	for idx, char := range input {
		if idx != 0 {
			if rune(input[idx-1]) == char {
				count++
			} else {
				count, res = 1, res+encode(count, idx, input)
			}
		}
	}
	if input != "" {
		res += encode(count, len(input), input)
	}

	return res
}

func RunLengthDecode(input string) string {
	count := 1
	var stringCount, res string
	for _, char := range input {
		if _, err := strconv.Atoi(string(char)); err == nil {
			stringCount += string(char)
		} else {
			if stringCount != "" {
				count, _ = strconv.Atoi(stringCount)
			}
			for j := 0; j < count; j++ {
				res += string(char)
			}
			count = 1
			stringCount = ""
		}
	}
	return res
}

func encode(count, i int, s string) string {
	if count > 1 {
		return fmt.Sprintf("%d%c", count, s[i-1])
	}
	return fmt.Sprintf("%c", s[i-1])
}
