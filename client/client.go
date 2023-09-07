package client

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/gorilla/websocket"
)

func Client(ip string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your nickname: ")
	nickname, _ := reader.ReadString('\n')

	// Remove newline character from the input
	nickname = nickname[:len(nickname)-1]

	serverAddr := "ws://" + ip + "/ws" // Change to your server's address
	fmt.Println(serverAddr)

	conn, _, err := websocket.DefaultDialer.Dial(serverAddr, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Send the player's nickname to the server
	err = conn.WriteMessage(websocket.TextMessage, []byte(nickname))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the server.")
	fmt.Println("You can now send messages to the server.")

	// Start a goroutine to read and display messages from the server
	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Fatal(err)
				return
			}
			fmt.Println(string(message))
		}
	}()

	// Read user input and send it to the server
	for {
		message, _ := reader.ReadString('\n')

		// Remove newline character from the input
		message = message[:len(message)-1]

		if message == "exit" {
			break
		}

		// Send the message to the server
		err := conn.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			log.Fatal(err)
		}
	}
}
