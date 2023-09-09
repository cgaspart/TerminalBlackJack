package blackjack

import (
	"encoding/json"
	"errors"
	"fmt"

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

func (p *Player) PrintHand(me bool) {
	HLColor := utils.HL_YELLO

	if me {
		HLColor = utils.HL_BLUE
	}
	val1, val2 := CardValue(p.Hand)
	fmt.Printf(`
	
  %s %s %s
value: %d`, HLColor, p.Name, utils.RESET, val1)
	if val2 != 0 {
		fmt.Print("/", val2)
	}

	PrintCards(p.Hand)

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
