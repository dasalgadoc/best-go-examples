package main

type AveragerActor struct {
	BaseActor
	bound float32
}

func NewAveragerActor(role string, gameUpperBound int) *AveragerActor {
	return &AveragerActor{
		BaseActor: BaseActor{
			name:           "Averager",
			role:           role,
			gameUpperBound: gameUpperBound,
			quantityEarned: 0,
			gamesPlayed:    []Game{},
		},
		bound: 0.05,
	}
}

func (a *AveragerActor) Propose(gameNumber int) (quantity int) {
	if gameNumber == 0 {
		return a.gameUpperBound / 2
	}

	sum := 0
	for _, proposal := range a.gamesPlayed {
		sum += proposal.Quantity
	}

	return sum / len(a.gamesPlayed)
}

func (a *AveragerActor) Accept(gameNumber int, partnerName string, quantity int) bool {
	if gameNumber == 0 {
		return quantity >= a.gameUpperBound/2
	}

	sum := 0
	for _, proposal := range a.gamesPlayed {
		sum += proposal.Quantity
	}

	avg := sum / len(a.gamesPlayed)

	return quantity >= (avg - int(float32(avg)*a.bound))
}
