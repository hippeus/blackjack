package blackjack

import (
	"errors"
	"log"

	"github.com/hippeus/deck"
)

type state uint8

const (
	statePlayerHand state = iota
	stateDealerHand
	stateHandOver
)

type Game struct {
	state    state
	shuffled bool
	deck     []deck.Card
	houseAI  AI

	player []deck.Card
	dealer []deck.Card
}

func New() Game {
	return Game{
		state:   statePlayerHand,
		deck:    nil,
		houseAI: dealerAI{},
	}
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
	case statePlayerHand:
		return &g.player
	case stateDealerHand:
		return &g.dealer
	default:
		// forbidden state
		return nil
	}
}

func (g *Game) Play(player AI) int {
	if !g.shuffled || g.deck == nil {
		g.deck = deck.New(deck.Deck(3), deck.Shuffle)
		g.shuffled = true
	}
	deal(g)
	for g.state == statePlayerHand {
		// TODO(hippeus): protect data by working on copy
		mv := player.Play(g.player, g.dealer[0])
		err := mv(g)
		if err != nil {
			log.Println(err, g.player)
			break
		}
	}
	for g.state == stateDealerHand {
		mv := g.houseAI.Play(g.dealer, g.dealer[0])
		err := mv(g)
		if err != nil {
			log.Println(err, g.dealer)
			break
		}
	}
	player.Result(g.player, g.dealer)
	return 0
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
