package blackjack

import (
	"fmt"

	"github.com/hippeus/deck"
)

type AI interface {
	Bet() int
	Play(hand []deck.Card, dealer deck.Card) Move
	Result(hand []deck.Card, dealHand []deck.Card)
}

func HumanPlayer() AI {
	return humanAI{}
}

type humanAI struct{}

func (h humanAI) Bet() int {
	return 1
}
func (h humanAI) Play(hand []deck.Card, dealer deck.Card) Move {
	var input string
	for {
		fmt.Printf("Player's hand: %s.\n", hand)
		fmt.Printf("Dealer's hand: %s.\n", dealer)
		fmt.Println("What's your move? Do you (h)it or (s)tand?")
		fmt.Scanf("%s", &input)
		switch input {
		case "h":
			return MoveHit
		case "s":
			return MoveStand
		default:
			fmt.Println("Invalid user input. Pick between (h)it and (s)tand")
		}
	}
}

func (h humanAI) Result(hand, dealer []deck.Card) {
	pScore := Score(hand...)
	dScore := Score(dealer...)

	fmt.Printf("Player's hand: %s.\nScored: %d\n", hand, pScore)
	fmt.Printf("Dealer's hand: %s.\nScored: %d\n", dealer, dScore)

	switch {
	case pScore > 21:
		fmt.Println("You busted")
	case dScore > 21:
		fmt.Println("Dealer busted")
	case pScore > dScore:
		fmt.Println("You win!")
	case dScore > pScore:
		fmt.Println("Dealer wins!")
	case dScore == pScore:
		fmt.Println("Draw")
	}
}

type dealerAI struct{}

func (ai dealerAI) Bet() int {
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
