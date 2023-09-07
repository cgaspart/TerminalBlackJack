package blackjack

import (
	"errors"

	"github.com/gorilla/websocket"
)

var (
	ErrCannotBetMore = errors.New("cannot bet more than balance")
)

type Player struct {
	Conn     *websocket.Conn
	Nickname string
	Cards    []Card
	Bet      float32
	Balance  float32
	Ready    bool
}

func NewPlayer(nick string, balance float32, con *websocket.Conn) *Player {
	player := &Player{
		Conn:     con,
		Nickname: nick,
		Bet:      0,
		Balance:  balance,
		Ready:    false,
	}

	return player
}

func (p *Player) PlayerReady() {
	p.Ready = true
}

func (p *Player) Betting(amount float32) error {
	if amount > p.Balance {
		return ErrCannotBetMore
	}

	p.Bet = amount
	return nil
}
