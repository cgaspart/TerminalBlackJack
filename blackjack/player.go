package blackjack

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/cgaspart/blackjack/utils"
	"github.com/gorilla/websocket"
)

var (
	ErrCannotBetMore = errors.New("cannot bet more than balance")
)

type Player struct {
	Conn    *websocket.Conn `json:"-"`
	Name    string          `json:"name"`
	Hand    []Card          `json:"hand"`
	Bet     float32         `json:"bet"`
	Balance float32         `json:"balance"`
	Ready   bool            `json:"ready"`
}

func NewPlayer(nick string, balance float32, con *websocket.Conn) *Player {
	player := &Player{
		Conn:    con,
		Name:    nick,
		Bet:     0,
		Balance: balance,
		Ready:   false,
	}

	return player
}

func (p *Player) Betting(amount float32) error {

	if amount > p.Balance {
		return ErrCannotBetMore
	}

	p.Bet = amount
	p.Balance = p.Balance - p.Bet
	return nil
}

func (p *Player) PrintHand() {
	val1, val2 := CardValue(p.Hand)
	fmt.Printf(`
	
  %s %s %s
value: %d`, utils.HIGHLIGHT, p.Name, utils.RESET, val1)
	if val2 != 0 {
		fmt.Print("/", val2)
	}

	maxRows := 0
	for _, card := range p.Hand {
		lines := strings.Split(card.CardArt, "\n")
		if len(lines) > maxRows {
			maxRows = len(lines)
		}
	}

	grid := make([]string, maxRows)

	// Populate the grid with the card arts
	for _, card := range p.Hand {
		lines := strings.Split(card.CardArt, "\n")
		for i, line := range lines {
			grid[i] += line
		}
	}

	for _, row := range grid {
		fmt.Println(row)
	}

	fmt.Printf(`
.-------------------.
| %s Bet: %s%.2f%s     |
| %s Balance: %s%.2f%s |
'-------------------'
`, p.Name, utils.BLUE, p.Bet, utils.RESET, p.Name, utils.GREEN, p.Balance, utils.RESET)
}

func (p *Player) SendPlayer() error {
	message := utils.Data{
		Type: utils.PLAYER,
		Data: p,
	}

	return utils.SendData(p.Conn, message)
}

func GetPlayer(data []byte) (*Player, error) {
	player := Player{}

	if err := json.Unmarshal(data, &player); err != nil {
		return nil, err
	}

	return &player, nil
}
