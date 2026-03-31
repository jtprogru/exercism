package blackjack

var Cards = map[string]int{
	"ace":   11,
	"eight": 8,
	"two":   2,
	"nine":  9,
	"three": 3,
	"ten":   10,
	"four":  4,
	"jack":  10,
	"five":  5,
	"queen": 10,
	"six":   6,
	"king":  10,
	"seven": 7,
	"other": 0,
}

var Opt = map[string]string{
	"Stand":             "S",
	"Hit":               "H",
	"Split":             "P",
	"Automatically win": "W",
}

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	return Cards[card]
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	handScore := ParseCard(card1) + ParseCard(card2)
	dealerScore := ParseCard(dealerCard)
	switch {
	case handScore == 22:
		return "P"
	case handScore == 21:
		if dealerScore < 10 {
			return "W"
		} else {
			return "S"
		}
	case handScore >= 17 || handScore >= 12 && dealerScore < 7:
		return "S"
	default:
		return "H"
	}
}
