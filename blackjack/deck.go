package blackjack

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Card struct {
	CardArt string
	Suit    string
	Rank    string
}

type Deck struct {
	Cards []Card
}

func NewDeck() *Deck {
	deck := &Deck{}
	suits := []string{"♥️", "♦️", "♣️", "♠️"}
	ranks := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}

	for _, suit := range suits {
		for _, rank := range ranks {
			card := Card{Suit: suit, Rank: rank}
			card.CardArt = fmt.Sprintf(`
		  .------.
		  |%s    %s|
		  |       |
		  |   %s  %s|
		  '------'`, rank, suit, suit, rank)
			deck.Cards = append(deck.Cards, card)
		}
	}

	return deck
}

func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
}

func (d *Deck) Deal() Card {
	if len(d.Cards) == 0 {
		return Card{} // Return an empty card if the deck is empty.
	}

	// Remove and return the first card from the deck.
	card := d.Cards[0]
	d.Cards = d.Cards[1:]
	return card
}

func printHand(hand []Card) string {
	var message string

	for _, card := range hand {
		message = message + fmt.Sprintln(card.CardArt)
	}

	return message
}

func cardValue(deck []Card) (int, int) {
	var value int
	ace := false

	for _, card := range deck {
		num, err := strconv.Atoi(card.Rank)
		if err != nil {
			if card.Rank == "A" {
				ace = true
			} else {
				value += 10
			}
		} else {
			value += num
		}
	}

	if ace {
		return value + 1, value + 11
	}

	return value, 0
}
