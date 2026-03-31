package allergies

import "math"

var allergens = []string{"eggs", "peanuts", "shellfish", "strawberries", "tomatoes", "chocolate", "pollen", "cats"}

func Allergies(allergies uint) []string {
	var res []string
	for _, v := range allergens {
		if AllergicTo(allergies, v) {
			res = append(res, v)
		}
	}
	return res
}

func AllergicTo(allergies uint, allergen string) bool {
	index := indexOf(allergens, allergen)
	res := uint(math.Pow(2.0, float64(index)))
	return (allergies & res) > 0
}

func indexOf(slice []string, allergen string) int {
	for idx, v := range slice {
		if v == allergen {
			return idx
		}
	}
	return -1
}
