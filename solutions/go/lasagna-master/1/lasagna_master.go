package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, minute int) int {
	if minute == 0 {
		return len(layers) * 2
	} else {
		return len(layers) * minute
	}
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	var noodles int
	var sauce float64
	for _, v := range layers {
		if v == "sauce" {
			sauce += 0.2
		} else if v == "noodles" {
			noodles += 50
		}
	}

	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList, myList []string) {
	myList[len(myList)-1] = friendList[len(friendList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(inputList []float64, scale int) []float64 {
	var scaled []float64
	for _, v := range inputList {
		scaled = append(scaled, (v/2)*float64(scale))
	}
	return scaled
}
