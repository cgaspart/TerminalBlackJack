package blackjack

import (
	"encoding/json"
	"fmt"

	"github.com/cgaspart/blackjack/utils"
	"github.com/gorilla/websocket"
)

type Game struct {
	Deck       *Deck     `json:"-"`
	DealerHand []Card    `json:"dealer_hand"`
	Players    []*Player `json:"player_list"`
}

func NewGame() *Game {
	game := &Game{}
	game.Deck = NewDeck()
	game.Deck.Shuffle()

	game.DealerHand = append(game.DealerHand, game.Deck.Deal())

	return game
}

func (g *Game) InitBet(player *Player) {
	for {
		message := fmt.Sprintf("Enter a bet amount\nYour blance is: %.2f", player.Balance)

		player.Conn.WriteMessage(websocket.TextMessage, []byte(message))

		betAmount, err := utils.GetMessageInt(player.Conn)
		if err != nil {
			if err == utils.ErrWrongNumber {
				player.Conn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
				continue
			}
			break
		}

		player.Betting(float32(betAmount))
		break
	}

}

func (g *Game) PrintDealerHand() {
	val1, val2 := CardValue(g.DealerHand)
	fmt.Printf(`
	
%s DEALER %s
VALUE: %d`, utils.HL_GREEN, utils.RESET, val1)
	if val2 != 0 {
		fmt.Print("/", val2)
	}

	PrintCards(g.DealerHand)
}

func (g *Game) PrintGame(currentPlayer *Player) {
	g.PrintDealerHand()

	for _, player := range g.Players {
		if player.Name != currentPlayer.Name {
			player.PrintHand(false)
		}
	}
	currentPlayer.PrintHand(true)
}

func (g *Game) SendGame() {
	message := utils.Data{
		Type: utils.GAME,
		Data: g,
	}

	for _, player := range g.Players {
		utils.SendData(player.Conn, message)
	}
}

func (g *Game) AddPlayer(player *Player) {
	g.InitBet(player)

	g.Players = append(g.Players, player)

	player.Hand = append(player.Hand, g.Deck.Deal())
	player.Hand = append(player.Hand, g.Deck.Deal())
}

func GetGame(data []byte) (*Game, error) {
	game := Game{}

	if err := json.Unmarshal(data, &game); err != nil {
		return nil, err
	}

	return &game, nil
}
