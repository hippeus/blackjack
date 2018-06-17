package main

import (
	"fmt"

	"github.com/hippeus/deck"
)

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}
func main() {
	cards := deck.New(deck.Deck(3), deck.Shuffle)
	var card deck.Card
	var player, dealer Hand

	for i := 0; i < 2; i++ {
		for _, hand := range []*Hand{&player, &dealer} {
			card, cards = draw(cards)
			*hand = append(*hand, card)
		}
	}

	var input string
	for input != "s" {
		fmt.Printf("Player's hand: %s.\n", player)
		fmt.Printf("Dealer's hand: %s.\n", dealer.dealerHand())
		fmt.Println("What's your move? Do you (h)it or (s)tand?")
		fmt.Scanf("%s", &input)
		switch input {
		case "h":
			card, cards = draw(cards)
			player = append(player, card)
		case "s":
		default:
			fmt.Println("Invalid user input. Pick between (h)it and (s)tand")
		}
	}

	pScore := player.Score()
	dScore := dealer.Score()

	for dScore < 16 || (dScore == 17 && dealer.minScore() != 17) {
		card, cards = draw(cards)
		dealer = append(dealer, card)
		dScore = dealer.Score()
	}

	//final
	fmt.Printf("Player's hand: %s.\nScored: %d\n", player, pScore)
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
