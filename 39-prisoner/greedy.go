package main

type GreedyActor struct {
	BaseActor
}

func NewGreedyActor(role string, gameUpperBound int) *GreedyActor {
	return &GreedyActor{
		BaseActor: BaseActor{
			name:           "Greedy",
			role:           role,
			gameUpperBound: gameUpperBound,
			quantityEarned: 0,
			gamesPlayed:    []Game{},
		},
	}
}

func (g *GreedyActor) Propose(gameNumber int) (quantity int) {
	return g.gameUpperBound / 10
}

func (g *GreedyActor) Accept(gameNumber int, partnerName string, quantity int) bool {
	return quantity > 2*g.gameUpperBound/3
}
