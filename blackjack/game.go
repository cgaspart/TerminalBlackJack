package blackjack

import (
	"fmt"

	"github.com/cgaspart/blackjack/utils"
	"github.com/gorilla/websocket"
)

type Game struct {
	Deck       *Deck
	DealerHand []Card
	Players    []*Player
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
		message := fmt.Sprintf("Enter a bet amount\n Your blance is: %.2f", player.Balance)

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

func (g *Game) AddPlayer(player *Player) {
	g.InitBet(player)

	g.Players = append(g.Players, player)

	player.Cards = append(player.Cards, g.Deck.Deal())
	player.Cards = append(player.Cards, g.Deck.Deal())
}

/*
func (g *Game) GetDealerHand() string {
	message := fmt.Sprintln("Dealer hand:")

}

func (g *Game) GetHands() string {

	for _, player := range g.Players {

	}
}
*/
