package blackjack

// func Shuffle(gs GameState) GameState {
// 	ret := clone(gs)
// 	ret.deck = deck.New(deck.Deck(3), deck.Shuffle)
// 	return ret
// }

// func Deal(gs GameState) GameState {
// 	var card deck.Card
// 	ret := clone(gs)
// 	for _, p := range []*Hand{&ret.player, &ret.dealer} {
// 		for i := 0; i < 2; i++ {
// 			card, ret.deck = draw(ret.deck)
// 			*p = append(*p, card)
// 		}
// 	}
// 	ret.state = StatePlayerHand
// 	return ret
// }

// func Hit(gs GameState) GameState {
// 	ret := clone(gs)
// 	var card deck.Card
// 	card, ret.deck = draw(ret.deck)
// 	player := ret.CurrentPlayer()
// 	*player = append(*player, card)
// 	if player.Score() > 21 {
// 		return Stand(ret)
// 	}
// 	return ret
// }

// func Stand(gs GameState) GameState {
// 	ret := clone(gs)
// 	ret.state++
// 	return ret
// }

// func (gs *GameState) CurrentPlayer() *Hand {
// 	var cur *Hand
// 	switch gs.state {
// 	case StatePlayerHand:
// 		cur = &gs.player
// 	case StateDealerHand:
// 		cur = &gs.dealer
// 	default:
// 		fmt.Println("forbidden state")
// 	}
// 	return cur
// }

// func Tally(gs GameState) GameState {
// 	ret := clone(gs)

// 	pScore := ret.player.Score()
// 	dScore := ret.dealer.Score()

// 	fmt.Printf("Player's hand: %s.\nScored: %d\n", ret.player, pScore)
// 	fmt.Printf("Dealer's hand: %s.\nScored: %d\n", ret.dealer, dScore)

// 	switch {
// 	case pScore > 21:
// 		fmt.Println("You busted")
// 	case dScore > 21:
// 		fmt.Println("Dealer busted")
// 	case pScore > dScore:
// 		fmt.Println("You win!")
// 	case dScore > pScore:
// 		fmt.Println("Dealer wins!")
// 	case dScore == pScore:
// 		fmt.Println("Draw")
// 	}
// 	ret.dealer, ret.player = nil, nil
// 	ret.state = StatePlayerHand
// 	return ret
// }
// func clone(gs GameState) GameState {
// 	return GameState{
// 		state:  gs.state,
// 		deck:   append(gs.deck[:0:0], gs.deck...),
// 		player: append(gs.player[:0:0], gs.player...),
// 		dealer: append(gs.dealer[:0:0], gs.dealer...),
// 	}
// }
