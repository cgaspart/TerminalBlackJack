package utils

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gorilla/websocket"
)

var (
	ErrWrongNumber = errors.New("invalid number")
)

const (
	PLAYER = "PLAYER"
	GAME   = "GAME"

	// COLOR
	RESET = "\033[0m"

	RED       = "\033[31m"
	GREEN     = "\033[32m"
	YELLOW    = "\033[33m"
	BLUE      = "\033[34m"
	HIGHLIGHT = "\033[43;30m"
)

type Generic struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

type Data struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

func SendData(conn *websocket.Conn, message Data) error {
	data, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
		return err
	}

	err = conn.WriteMessage(websocket.BinaryMessage, data)
	if err != nil {
		return err
	}
	return nil
}

func GetMessageInt(conn *websocket.Conn) (int, error) {
	_, p, err := conn.ReadMessage()
	if err != nil {
		log.Println(err)
		return 0, err
	}

	betStr := string(p)
	number, err := strconv.Atoi(betStr)
	if err != nil {
		log.Println(err)
		return 0, ErrWrongNumber
	}

	return number, nil
}

func GetMessageString(conn *websocket.Conn) (string, error) {
	_, p, err := conn.ReadMessage()
	if err != nil {
		log.Println(err)
		return "", err
	}
	message := string(p)

	return message, nil
}

func GetUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	return input[:len(input)-1]
}
