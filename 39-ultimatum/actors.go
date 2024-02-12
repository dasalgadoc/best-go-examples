package main

type Actor interface {
	GetName() string
	Propose(gameNumber int) (quantity int)
	Accept(gameNumber int, partnerName string, quantity int) bool
	RegisterGame(game Game)
	GetOverall() int
	Restart()
}

type BaseActor struct {
	name           string
	role           string
	gameUpperBound int
	quantityEarned int
	gamesPlayed    []Game
}

func (b *BaseActor) GetName() string {
	return b.name
}

func (b *BaseActor) GetRole() string {
	return b.role
}

func (b *BaseActor) RegisterGame(game Game) {
	b.gamesPlayed = append(b.gamesPlayed, game)
	if game.Accepted {
		if b.role == "proposer" {
			b.quantityEarned += b.gameUpperBound - game.Quantity
		} else {
			b.quantityEarned += game.Quantity
		}
	}
}

func (b *BaseActor) GetOverall() int {
	return b.quantityEarned
}

func (b *BaseActor) GetGamesPlayed() []Game {
	return b.gamesPlayed
}

func (b *BaseActor) Propose(gameNumber int) (quantity int) {
	return
}

func (b *BaseActor) Accept(gameNumber int, partnerName string, quantity int) bool {
	return false
}

func (b *BaseActor) Restart() {
	b.quantityEarned = 0
	b.gamesPlayed = []Game{}
}
