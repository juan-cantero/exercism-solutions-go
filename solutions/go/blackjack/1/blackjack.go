package blackjack

var cards = map[string]int{
	"ace":   11,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"ten":   10,
	"jack":  10,
	"queen": 10,
	"king":  10,
}

const (
	ActionSplit = "P" // Split
	ActionWin   = "W" // Automatically win
	ActionStand = "S" // Stand
	ActionHit   = "H" // Hit
)

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	return cards[card]
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {

	pCard1 := ParseCard(card1)
	pCard2 := ParseCard(card2)
	pDealerCard := ParseCard(dealerCard)
	playerScore := pCard1 + pCard2

	switch {
	case card1 == card2 && card1 == "ace":
		return ActionSplit
	case playerScore == 21:
		if pDealerCard < 10 {
			return ActionWin
		}
		return ActionStand
	case playerScore >= 17 && playerScore <= 20:
		return ActionStand

	case playerScore >= 12 && playerScore < 17 && pDealerCard < 7:
		return ActionStand

	default:
		return ActionHit
	}
}
