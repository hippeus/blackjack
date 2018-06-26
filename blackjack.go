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

var (
	errBusted = errors.New("BUSTED")
)

type Options struct {
	Decks           uint
	Rounds          uint
	Wager           uint
	BlackJackPayout float64
}

type Game struct {
	state           state
	nRounds         uint
	blackJackPayout float64

	deck     []deck.Card
	nDecks   uint
	shuffled bool

	player  []deck.Card
	balance float64
	bet     uint

	dealer  []deck.Card
	houseAI AI
}

func New(opts Options) Game {
	var g = Game{
		state:           statePlayerTurn,
		houseAI:         dealerAI{},
		nRounds:         1,
		blackJackPayout: 1.5,
		bet:             1,
		nDecks:          3,
		deck:            nil,
		player:          nil,
		dealer:          nil,
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
	if opts.BlackJackPayout != 0 {
		g.blackJackPayout = opts.BlackJackPayout
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
		return errBusted
	}
	return nil
}

func MoveStand(g *Game) error {
	g.state++
	return nil
}

func MoveDouble(g *Game) error {
	if len(g.player) != 2 {
		return errors.New("You can only double on a hand with 2 cards")
	}
	g.bet += g.bet
	var c deck.Card
	c, g.deck = draw(g.deck)

	currHand := currentPlayer(g)
	*currHand = append(*currHand, c)
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
	for i := uint(0); i < g.nRounds; i++ {
		fmt.Printf("*** Round %d ***\n", i+1)
		if len(g.deck) <= int(g.nDecks*52)/3 {
			g.shuffled = false
		}
		if !g.shuffled || g.deck == nil {
			g.deck = deck.New(deck.Deck(3), deck.Shuffle)
			g.shuffled = true
		}
		bet(g, player)
		deal(g)
		// g.player[0] = deck.Card{Rank: deck.Ten}
		// g.player[1] = deck.Card{Rank: deck.Nine}
		// g.dealer[0] = deck.Card{Rank: deck.Ace}
		// g.dealer[1] = deck.Card{Rank: deck.Jack}
		dealerBJ := blackjack(g.dealer[0], g.dealer[1])
		if dealerBJ {
			endRound(g, player)
			continue
		}
		for g.state == statePlayerTurn {
			hand := make([]deck.Card, len(g.player))
			copy(hand, g.player)
			mv := player.Play(hand, g.dealer[0])
			err := mv(g)
			switch err {
			case errBusted:
				g.state = stateHandOver
			case nil:
				// noop
			default:
				fmt.Println(err)
			}
		}
		for g.state == stateDealerTurn {
			hand := make([]deck.Card, len(g.dealer))
			copy(hand, g.dealer)
			mv := g.houseAI.Play(hand, g.dealer[0])
			err := mv(g)
			switch err {
			case errBusted:
				g.state = stateHandOver
			case nil:
				// noop
			default:
				fmt.Println(err)
			}
		}
		endRound(g, player)
	}
	return int(g.balance)
}

func bet(g *Game, ai AI) {
	g.bet = ai.Bet()
}

func endRound(g *Game, ai AI) {
	dealerBJ := blackjack(g.dealer[0], g.dealer[1])
	playerBJ := blackjack(g.player[0], g.player[1])
	pScore := Score(g.player...)
	dScore := Score(g.dealer...)
	roundWin := g.evalResult(pScore, dScore, playerBJ, dealerBJ)
	g.balance += roundWin

	ai.Result(g.player, g.dealer)
	g.player, g.dealer = nil, nil
	g.state = statePlayerTurn
}

func (g *Game) evalResult(pScore, dScore int, playerBJ, dealerBJ bool) float64 {
	var win float64
	switch {
	case playerBJ && dealerBJ:
		fmt.Println("Draw")
		win = 0
	case playerBJ:
		win = float64(g.bet) * g.blackJackPayout
	case dealerBJ:
		win = -float64(g.bet)
	case pScore > 21:
		fmt.Println("You busted")
		win -= float64(g.bet)
	case dScore > 21:
		fmt.Println("Dealer busted")
		win += float64(g.bet)
	case pScore > dScore:
		fmt.Println("You win!")
		win += float64(g.bet)
	case dScore > pScore:
		fmt.Println("Dealer wins!")
		win -= float64(g.bet)
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

func blackjack(first, second deck.Card) bool {
	if !(Score(first, second) == 21) {
		return false
	}
	return true
}
