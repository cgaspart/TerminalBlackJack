package server

import (
	"fmt"

	"github.com/cgaspart/blackjack/blackjack"
	"github.com/cgaspart/blackjack/utils"
)

func countPlayerNotReady() int {
	notReady := 0

	for _, player := range connections {
		if !player.Ready {
			notReady++
		}
	}

	return notReady
}

func JoinWaitingRoom(player *blackjack.Player) {
	for {
		game := blackjack.NewGame()

		message, err := utils.GetMessageString(player.Conn)
		if err != nil {
			break
		}

		switch message {
		case "ready":
			player.Ready = true

			game.AddPlayer(player)

			notReady := countPlayerNotReady()
			if notReady == 0 {
				message := "All players ready\nLaunching a new game..."
				broadcast(message)

			}
			message = fmt.Sprintf("Player %s is ready\nWaiting for %d more player", player.Name, notReady)
			broadcast(message)
		}
	}
}
