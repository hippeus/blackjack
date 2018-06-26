package blackjack

import "github.com/hippeus/deck"

func Score(cards ...deck.Card) int {
	sum := minScore(cards...)
	if sum > 11 {
		return sum
	}
	for _, card := range cards {
		if card.Rank == deck.Ace && sum <= 11 {
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
