package client

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/cgaspart/blackjack/blackjack"
	"github.com/cgaspart/blackjack/utils"
	"github.com/gorilla/websocket"
)

var (
	conn   *websocket.Conn
	Player *blackjack.Player
	Game   *blackjack.Game
)

func login() {
	var err error
	ip := utils.GetUserInput("Enter server ip and port (localhost:888): ")
	userName := utils.GetUserInput("Enter your nickname: ")

	serverAddr := "ws://" + ip + "/ws"

	conn, _, err = websocket.DefaultDialer.Dial(serverAddr, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = conn.WriteMessage(websocket.TextMessage, []byte(userName))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the server.")
}

func handleServerData(message []byte) error {
	var err error
	srvData := utils.Generic{}

	if err := json.Unmarshal(message, &srvData); err != nil {
		return err
	}

	switch srvData.Type {
	case utils.PLAYER:
		Player, err = blackjack.GetPlayer(srvData.Data)
		if err != nil {
			log.Fatal(err)
			return err
		}
	case utils.GAME:
		Game, err = blackjack.GetGame(srvData.Data)
		if err != nil {
			log.Fatal(err)
			return err
		}
		Game.PrintGame(Player)
	}
	return nil
}

func Client() {
	defer conn.Close()

	login()

	go func() {
		for {
			messageType, message, err := conn.ReadMessage()
			if err != nil {
				log.Fatal(err)
				return
			}
			switch messageType {
			case websocket.TextMessage:
				fmt.Println(string(message))
			case websocket.BinaryMessage:
				if err := handleServerData(message); err != nil {
					log.Fatal(err)
				}
			}

		}
	}()

loop:
	for {
		command := utils.GetUserInput("")

		switch command {
		case "exit":
			break loop
		case "/me":
			if Player != nil {
				Player.PrintHand(true)
			} else {
				fmt.Println("no hand found")
			}
		case "/dealer":
			if Game != nil {
				Game.PrintDealerHand()
			} else {
				fmt.Println("no game found")
			}
		}
		err := conn.WriteMessage(websocket.TextMessage, []byte(command))
		if err != nil {
			log.Fatal(err)
		}
	}
}
