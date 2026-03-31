package armstrongnumbers

func IsNumber(n int) bool {
	if n < 0 {
		return false
	}
	if n < 10 {
		return true
	}

	numDigits := 0
	for temp := n; temp > 0; temp /= 10 {
		numDigits++
	}

	sum := 0
	for temp := n; temp > 0; temp /= 10 {
		digit := temp % 10
		pow := 1
		for range numDigits {
			pow *= digit
		}
		sum += pow
	}

	return sum == n
}
