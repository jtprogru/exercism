package dndcharacter

import (
	"math"
	"math/rand/v2"
	"sort"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	return int(math.Floor(float64(score-10) / 2))
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	dice := []int{roll(), roll(), roll(), roll()}
	sort.Ints(dice) // сортируем по возрастанию: наименьший окажется в dice[0]
	// Отбрасываем dice[0] (минимальный) и складываем три больших.
	return dice[1] + dice[2] + dice[3]
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	c := Character{
		Strength:     Ability(),
		Dexterity:    Ability(),
		Constitution: Ability(),
		Intelligence: Ability(),
		Wisdom:       Ability(),
		Charisma:     Ability(),
	}
	// Хитпоинты = 10 + модификатор телосложения (Constitution).
	c.Hitpoints = 10 + Modifier(c.Constitution)
	return c
}

func roll() int {
	return rand.IntN(6) + 1
}
