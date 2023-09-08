package main

import (
	"fmt"
	"os"
	"time"

	"github.com/cgaspart/blackjack/client"
	"github.com/cgaspart/blackjack/server"
	"github.com/cgaspart/blackjack/utils"
)

func main() {
	fmt.Println("Welcome to Blackjack Terminal Game!")
	fmt.Println("Please choose an option:")
	fmt.Println("1. Join a Game")
	fmt.Println("2. Host a Game")
	fmt.Println("3. Quit")
	for {
		option := utils.GetUserInput("")

		switch option {
		case "1":
			client.Client()
		case "2":
			go server.RunSRV()
			time.Sleep(2 * time.Second)
			client.Client()
		case "3":
			os.Exit(0)
		default:
			fmt.Println("invalid option")
		}
	}

}
