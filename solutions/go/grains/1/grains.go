package grains

import "fmt"

func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, fmt.Errorf("Input[%d] invalid. Input should be between 1 and 64 (inclusive)", number)
	}
	return 1 << (uint16(number) - 1), nil
}

func Total() uint64 {
	return (1 << 64) - 1
}
