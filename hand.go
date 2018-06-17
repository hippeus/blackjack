package main

import (
	"strings"

	"github.com/hippeus/deck"
)

type Hand []deck.Card

func (h Hand) String() string {
	toString := make([]string, len(h))
	for i := range h {
		toString[i] = h[i].String()
	}
	return strings.Join(toString, ", ")
}

func (h Hand) dealerHand() string {
	return h[0].String() + ", ***HIDDEN***"
}

func (h Hand) Score() int {
	sum := h.minScore()
	if sum > 11 {
		return sum
	}
	for _, card := range h {
		if card.Rank == deck.Ace {
			sum += 10
		}
	}
	return sum
}

func (h Hand) minScore() int {
	var sum int
	for _, card := range h {
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
