package server

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func login(w http.ResponseWriter, r *http.Request) (*websocket.Conn, string) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return nil, ""
	}
	defer conn.Close()

	// Read the player's nickname
	_, nickname, err := conn.ReadMessage()
	if err != nil {
		log.Println(err)
		return nil, ""
	}

	playerName := string(nickname)

	return conn, playerName
}
