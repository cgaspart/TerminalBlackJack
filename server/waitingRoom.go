package server

import (
	"fmt"
	"log"
	"strconv"

	"github.com/cgaspart/blackjack/blackjack"
	"github.com/gorilla/websocket"
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
		_, p, err := player.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		message := string(p)
		switch message {
		case "ready":
			connections[nicknameStr].PlayerReady()
			for {
				message := fmt.Sprintf("Enter a bet amount\n Your blance is: %.2f", connections[nicknameStr].Balance)

				connections[nicknameStr].Conn.WriteMessage(websocket.TextMessage, []byte(message))

				_, p, err := conn.ReadMessage()
				if err != nil {
					log.Println(err)
					break
				}
				betStr := string(p)
				num, err := strconv.Atoi(betStr)
				if err != nil {
					connections[nicknameStr].Conn.WriteMessage(websocket.TextMessage, []byte("ENTER A VALID NUMBER"))
				}
				connections[nicknameStr].Betting(float32(num))
				message = fmt.Sprintf("Player %s is ready\nWaiting for %d more player", nicknameStr, countPlayNotReady())
				broadcast(message)
				break
			}
		}
		log.Printf("Received message from %s: %s\n", nicknameStr, message)

		if !inGame {
			allPlayerReady := true

			for _, player := range connections {
				if !player.Ready {
					allPlayerReady = false
					break
				}
			}

			if allPlayerReady {
				message := "All players ready\nLaunching a new game..."
				broadcast(message)
				game = blackjack.NewGame()
				inGame = true

				for _, player := range connections {
					game.AddPlayer(player)
				}
			}
		}
	}
}
