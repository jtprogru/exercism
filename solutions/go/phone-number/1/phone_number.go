package phonenumber

import (
	"errors"
	"strings"
)

// allowedPunctuations — символы форматирования, которые допустимо встретить
// в номере и которые мы просто отбрасываем (скобки, дефисы, точки, пробелы, плюс).
const allowedPunctuations = " .+-()"

// Number очищает телефонный номер от форматирования, проверяет его на
// соответствие правилам NANP и возвращает канонические 10 цифр.
// При любом нарушении правил возвращает ошибку.
func Number(phoneNumber string) (string, error) {
	// builder эффективно накапливает только цифры (без лишних аллокаций).
	var builder strings.Builder

	for _, r := range phoneNumber {
		switch {
		case r >= '0' && r <= '9':
			// Цифра — оставляем.
			builder.WriteRune(r)
		case strings.ContainsRune(allowedPunctuations, r):
			// Разрешённый символ форматирования — игнорируем.
			continue
		default:
			// Буква или запрещённая пунктуация — номер невалиден.
			return "", errors.New("invalid character in phone number")
		}
	}

	digits := builder.String()

	// 11 цифр допустимы, только если это код страны "1" впереди — убираем его.
	if len(digits) == 11 && digits[0] == '1' {
		digits = digits[1:]
	}

	// После нормализации номер обязан состоять ровно из 10 цифр.
	if len(digits) != 10 {
		return "", errors.New("phone number must have 10 digits")
	}

	// Первая цифра area code (N) — от 2 до 9.
	if digits[0] < '2' || digits[0] > '9' {
		return "", errors.New("area code must start with a digit from 2 to 9")
	}

	// Первая цифра exchange code (N) — от 2 до 9.
	if digits[3] < '2' || digits[3] > '9' {
		return "", errors.New("exchange code must start with a digit from 2 to 9")
	}

	return digits, nil
}

// AreaCode возвращает трёхзначный код региона (первые три цифры номера).
func AreaCode(phoneNumber string) (string, error) {
	number, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return number[:3], nil
}

// Format возвращает номер в человекочитаемом виде "(NXX) NXX-XXXX".
func Format(phoneNumber string) (string, error) {
	number, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	// number[0:3] — area code, number[3:6] — exchange code, number[6:10] — subscriber.
	return "(" + number[:3] + ") " + number[3:6] + "-" + number[6:], nil
}
