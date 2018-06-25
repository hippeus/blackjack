package main

import (
	"fmt"

	"github.com/hippeus/blackjack"
)

func main() {
	opts := blackjack.Options{Rounds: 3}
	game := blackjack.New(opts)
	wins := game.Play(blackjack.HumanPlayer())
	fmt.Println(wins)
}
