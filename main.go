package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/cgaspart/blackjack/client"
	"github.com/cgaspart/blackjack/server"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Run 1) server 2) client: ")
	option, _ := reader.ReadString('\n')

	option = option[:len(option)-1]

	if option == "1" {
		server.RunSRV()
	} else {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("enter srv IP and port (localhost:888): ")
		ip, _ := reader.ReadString('\n')

		ip = ip[:len(ip)-1]

		client.Client(ip)
	}

	/*
		deck := NewDeck()
		deck.Shuffle()

		var dealer []Card
		var player []Card
		// Deal a couple of cards and print them as an example.
		for i := 0; i < 2; i++ {
			dealer = append(dealer, deck.Deal())
			player = append(player, deck.Deal())
		}

		fmt.Printf("Dealer Hand: %s %s\n", dealer[0].Rank, dealer[0].Suit)

		fmt.Println("Player Hand: ")
		for _, card := range player {
			fmt.Printf("%s %s ", card.Rank, card.Suit)
		}
		value, valueAce := cardValue(player)

		fmt.Println("Value ", value)
		if valueAce != 0 {
			fmt.Println(valueAce)
		}
	*/
}
