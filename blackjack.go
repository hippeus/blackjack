package blackjack

import (
	"errors"
	"fmt"

	"github.com/hippeus/deck"
)

type state uint8

const (
	statePlayerTurn state = iota
	stateDealerTurn
	stateHandOver
)

type Options struct {
	Decks  uint
	Rounds uint
	Wager  uint
}

type Game struct {
	state   state
	nRounds uint
	houseAI AI

	deck     []deck.Card
	nDecks   uint
	shuffled bool

	player []deck.Card
	bet    uint

	dealer []deck.Card
}

func New(opts Options) Game {
	var g = Game{
		state:   statePlayerTurn,
		houseAI: dealerAI{},
		nRounds: 1,
		bet:     1,
		nDecks:  3,
		deck:    nil,
		player:  nil,
		dealer:  nil,
	}
	if opts.Wager != 0 {
		g.bet = opts.Wager
	}
	if opts.Rounds != 0 {
		g.nRounds = opts.Rounds
	}
	if opts.Decks != 0 {
		g.nDecks = opts.Decks
	}
	return g
}

type Move func(*Game) error

func MoveHit(g *Game) error {
	var card deck.Card
	card, g.deck = draw(g.deck)
	player := currentPlayer(g)
	*player = append(*player, card)
	if Score(*player...) > 21 {
		return errors.New("BUSTED")
	}
	return nil
}

func MoveStand(g *Game) error {
	g.state++
	return nil
}

func currentPlayer(g *Game) *[]deck.Card {
	switch g.state {
	case statePlayerTurn:
		return &g.player
	case stateDealerTurn:
		return &g.dealer
	default:
		// forbidden state
		return nil
	}
}

func (g *Game) Play(player AI) int {
	var win int
	for i := uint(0); i < g.nRounds; i++ {
		fmt.Printf("Round: %d\n", i)
		if len(g.deck) <= int(g.nDecks*52)/3 {
			g.shuffled = false
		}
		if !g.shuffled || g.deck == nil {
			g.deck = deck.New(deck.Deck(3), deck.Shuffle)
			g.shuffled = true
		}
		deal(g)
		for g.state == statePlayerTurn {
			hand := make([]deck.Card, len(g.player))
			copy(hand, g.player)
			mv := player.Play(hand, g.dealer[0])
			err := mv(g)
			if err != nil {
				break
			}
		}
		for g.state == stateDealerTurn {
			hand := make([]deck.Card, len(g.dealer))
			copy(hand, g.dealer)
			mv := g.houseAI.Play(hand, g.dealer[0])
			err := mv(g)
			if err != nil {
				break
			}
		}
		pScore := Score(g.player...)
		dScore := Score(g.dealer...)

		player.Result(g.player, g.dealer)
		roundWin := g.evalResult(pScore, dScore)
		win += roundWin
		g.player, g.dealer = nil, nil
		g.state = statePlayerTurn
	}
	return win
}

func (g *Game) evalResult(pScore, dScore int) int {
	win := 0
	switch {
	case pScore > 21:
		fmt.Println("You busted")
		win -= int(g.bet)
	case dScore > 21:
		fmt.Println("Dealer busted")
		win += int(g.bet)
	case pScore > dScore:
		fmt.Println("You win!")
		win += int(g.bet)
	case dScore > pScore:
		fmt.Println("Dealer wins!")
		win -= int(g.bet)
	case dScore == pScore:
		fmt.Println("Draw")
		win = 0
	}
	return win
}
func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	card := cards[0]
	cards[0] = deck.Card{}
	return card, cards[1:]
}

func deal(g *Game) {
	g.player = make([]deck.Card, 0, 5)
	g.dealer = make([]deck.Card, 0, 5)
	var card deck.Card

	for i := 0; i < 2; i++ {
		for _, p := range []*[]deck.Card{&g.player, &g.dealer} {
			card, g.deck = draw(g.deck)
			*p = append(*p, card)
		}
	}
}
