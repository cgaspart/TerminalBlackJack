package client

import (
	"fmt"
	"log"

	"github.com/cgaspart/blackjack/utils"
	"github.com/gorilla/websocket"
)

var (
	conn *websocket.Conn
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

func Client() {
	defer conn.Close()

	login()

	// Start a goroutine to read and display messages from the server
	go func() {
		for {
			message := utils.GetServerMessage(conn)
			fmt.Println(message)
		}
	}()

	for {
		command := utils.GetUserInput("")

		if command == "exit" {
			break
		}

		// Send the message to the server
		err := conn.WriteMessage(websocket.TextMessage, []byte(command))
		if err != nil {
			log.Fatal(err)
		}
	}
}
