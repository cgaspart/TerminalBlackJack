package utils

import (
	"bufio"
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

func GetServerMessage(conn *websocket.Conn) string {
	_, message, err := conn.ReadMessage()
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return string(message)
}
