package blackjack

import (
	"fmt"

	"github.com/hippeus/deck"
)

type AI interface {
	Bet() uint
	Play(hand []deck.Card, dealer deck.Card) Move
	Result(hand []deck.Card, dealHand []deck.Card)
}

func HumanPlayer() AI {
	return humanAI{}
}

type humanAI struct{}

func (h humanAI) Bet() uint {
	fmt.Println("How much do you want to bet?")
	var bet uint
	if _, err := fmt.Scanf("%d\n", &bet); err != nil {
		return 1
	}
	return bet
}
func (h humanAI) Play(hand []deck.Card, dealer deck.Card) Move {
	var input string
	for {
		fmt.Printf("Player's hand: %s.\n", hand)
		fmt.Printf("Dealer's hand: %s.\n", dealer)
		if blackjack(hand[0], hand[1]) {
			return MoveStand
		}
		fmt.Println("What's your move? Do you (h)it, (d)ouble or (s)tand?")
		fmt.Scanf("%s", &input)
		switch input {
		case "h":
			return MoveHit
		case "d":
			return MoveDouble
		case "s":
			return MoveStand
		default:
			fmt.Println("Invalid user input")
		}
	}
}

func (h humanAI) Result(hand, dealer []deck.Card) {
	fmt.Printf("***FINAL HANDS***\n")
	fmt.Printf("Player: %s.\n", hand)
	fmt.Printf("Dealer: %s.\n", dealer)
	fmt.Printf("***¯\\_(ツ)_/¯***\n")
}

type dealerAI struct{}

func (ai dealerAI) Bet() uint {
	// noop
	return 0
}

func (ai dealerAI) Play(hand []deck.Card, dealer deck.Card) Move {
	res := Score(hand...)
	if res < 16 || (res == 17 && minScore(hand...) != 17) {
		return MoveHit
	}
	return MoveStand
}

func (ai dealerAI) Result(hand []deck.Card, dealHand []deck.Card) {
	// noop
}
