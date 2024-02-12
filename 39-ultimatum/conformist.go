package main

type ConformistActor struct {
	BaseActor
}

func NewConformistActor(role string, gameUpperBound int) *ConformistActor {
	return &ConformistActor{
		BaseActor: BaseActor{
			name:           "Conformist",
			role:           role,
			gameUpperBound: gameUpperBound,
			quantityEarned: 0,
			gamesPlayed:    []Game{},
		},
	}
}

func (c *ConformistActor) Propose(gameNumber int) (quantity int) {
	return 50
}

func (c *ConformistActor) Accept(gameNumber int, partnerName string, quantity int) bool {
	return true
}
