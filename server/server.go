package server

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/cgaspart/blackjack/blackjack"
	"github.com/gorilla/websocket"
)

const (
	BALANCE_DEFAULT = 100
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	connections   = make(map[string]*blackjack.Player) // Map of WebSocket connections to player nicknames
	connectionsMu sync.Mutex                           // Mutex to protect the connections map
	inGame        = false
	game          *blackjack.Game
)

func handleConnection(w http.ResponseWriter, r *http.Request) {
	conn, playerName := login(w, r)

	defer conn.Close()

	log.Printf("New player connected with nickname: %s\n", playerName)

	// Lock the connections map while adding the new connection
	connectionsMu.Lock()
	connections[playerName] = blackjack.NewPlayer(playerName, BALANCE_DEFAULT, conn)
	connectionsMu.Unlock()

	// Send a message to all connected clients about the new player
	message := fmt.Sprintf("Player %s has joined the game. \nEnter 'ready' to launch the game", playerName)
	broadcast(message)

	// Main loop
	for {
		JoinWaitingRoom(connections[playerName])
	}

	// Remove the connection from the connections map when the client disconnects
	connectionsMu.Lock()
	delete(connections, playerName)
	connectionsMu.Unlock()

	message = fmt.Sprintf("Player %s has left the game.", playerName)
	broadcast(message)
}

func broadcast(message string) {
	connectionsMu.Lock()
	defer connectionsMu.Unlock()

	for _, player := range connections {
		err := player.Conn.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			log.Println(err)
		}
	}
}

func RunSRV() {
	http.HandleFunc("/ws", handleConnection)

	port := 888
	server := fmt.Sprintf("0.0.0.0:%d", port)
	fmt.Printf("WebSocket server is running on: %s\n", server)
	err := http.ListenAndServe(server, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
