package perfect

import "errors"

// Define the Classification type here.
type Classification string

const (
	ClassificationDeficient Classification = "deficient"
	ClassificationPerfect   Classification = "perfect"
	ClassificationAbundant  Classification = "abundant"
)

var ErrOnlyPositive = errors.New("positive number is required")

func Classify(n int64) (Classification, error) {
	if n <= 0 {
		return "", ErrOnlyPositive
	}
	var sum int64
	for i := int64(1); i < n; i++ {
		if n%i == 0 {
			if sum += i; sum > n {
				return ClassificationAbundant, nil
			}
		}
	}
	if sum == n {
		return ClassificationPerfect, nil
	} else if sum < n {
		return ClassificationDeficient, nil
	}
	return ClassificationAbundant, nil
}
