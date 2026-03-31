package romannumerals

import (
	"bytes"
	"fmt"
)

type arabicToRoman struct {
	arabic int
	roman  string
}

func ToRomanNumeral(input int) (string, error) {
	buf := bytes.NewBufferString("")

	if input <= 0 || input >= 3001 {
		return "", fmt.Errorf("the number %d is undefined in the roman numeral system", input)
	}

	mapp := []arabicToRoman{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	for _, m := range mapp {
		for input >= m.arabic {
			buf.WriteString(m.roman)
			input -= m.arabic
		}
	}

	return buf.String(), nil
}
