package server

import (
	"fmt"

	"github.com/cgaspart/blackjack/blackjack"
	"github.com/cgaspart/blackjack/utils"
)

func PlayerGameLoop(game *blackjack.Game) {
mainloop:
	for {
		for _, player := range game.Players {
		outerloop:
			for {
				game.SendGame()
				message, err := utils.GetMessageString(player.Conn)
				if err != nil {
					break
				}

				switch message {
				case "hit":
					playerBust, BJ := player.Hit(game)
					if BJ {
						break outerloop
					}
					if playerBust {
						message = fmt.Sprintf("Player %s bust", player.Name)
						broadcast(message)
						game.SendGame()
						break outerloop
					}
					message = fmt.Sprintf("Player %s hit", player.Name)
					broadcast(message)
					game.SendGame()
				case "stand":
					break outerloop
				}
			}
		}
		dealerBust := game.DealerTurn()
		if dealerBust {
			message := "Dealer bust"
			broadcast(message)
		}
		break mainloop
	}
}
