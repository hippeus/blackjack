package main

import (
	"fmt"
)

func main() {
	var gs GameState
	gs = Shuffle(gs)
	for i := 0; i < 3; i++ {
		gs = Deal(gs)
		var input string
		for gs.state == StatePlayerHand {
			fmt.Printf("Player's hand: %s.\n", gs.player)
			fmt.Printf("Dealer's hand: %s.\n", gs.dealer.dealerHand())
			fmt.Println("What's your move? Do you (h)it or (s)tand?")
			fmt.Scanf("%s", &input)
			switch input {
			case "h":
				gs = Hit(gs)
			case "s":
				gs = Stand(gs)
			default:
				fmt.Println("Invalid user input. Pick between (h)it and (s)tand")
			}
		}

		fmt.Println("Dealer TURN")
		for gs.state == StateDealerHand {
			for gs.dealer.Score() < 16 || (gs.dealer.Score() == 17 && gs.dealer.minScore() != 17) {
				gs = Hit(gs)
			}
			gs = Stand(gs)
		}
		gs = Tally(gs)
	}
}
