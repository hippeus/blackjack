package blackjack

import "github.com/hippeus/deck"

// type Hand []deck.Card

// func (h Hand) String() string {
// 	toString := make([]string, len(h))
// 	for i := range h {
// 		toString[i] = h[i].String()
// 	}
// 	return strings.Join(toString, ", ")
// }

// func (h Hand) dealerHand() string {
// 	return h[0].String() + ", ***HIDDEN***"
// }

func Score(cards ...deck.Card) int {
	sum := minScore(cards...)
	if sum > 11 {
		return sum
	}
	for _, card := range cards {
		if card.Rank == deck.Ace && sum < 11 {
			sum += 10
		}
	}
	return sum
}

func minScore(cards ...deck.Card) int {
	var sum int
	for _, card := range cards {
		switch card.Rank {
		case deck.King, deck.Queen, deck.Jack:
			sum += 10
		case deck.Ace:
			sum++
		default:
			sum += int(card.Rank)
		}
	}
	return sum
}
