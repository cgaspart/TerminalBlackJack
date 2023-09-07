package blackjack

type Game struct {
	Deck       *Deck
	DealerHand []Card
	Player     []*Player
}

func NewGame() *Game {
	game := &Game{}
	game.Deck = NewDeck()
	game.Deck.Shuffle()

	game.DealerHand = append(game.DealerHand, game.Deck.Deal())
	game.DealerHand = append(game.DealerHand, game.Deck.Deal())

	return game
}

func (g *Game) AddPlayer(player *Player) {
	g.Player = append(g.Player, player)

	player.Cards = append(player.Cards, g.Deck.Deal())
	player.Cards = append(player.Cards, g.Deck.Deal())
}
