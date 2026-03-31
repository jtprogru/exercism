package reverse

func Reverse(input string) string {
	if len(input) == 0 {
		return ""
	}
	var res string
	str := []rune(input)

	for i := len(str) - 1; i != -1; i-- {
		res += string(str[i])
	}
	return res
}
