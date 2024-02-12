package main

type JustifierActor struct {
	BaseActor
	midUpperPoint int
	tolerance     int
}

func NewJustifierActor(role string, gameUpperBound int) *JustifierActor {
	return &JustifierActor{
		BaseActor: BaseActor{
			name:           "Justifier",
			role:           role,
			gameUpperBound: gameUpperBound,
			quantityEarned: 0,
			gamesPlayed:    []Game{},
		},
		midUpperPoint: gameUpperBound / 2,
		tolerance:     5,
	}
}

func (j *JustifierActor) Propose(gameNumber int) (quantity int) {
	return j.midUpperPoint
}

func (j *JustifierActor) Accept(gameNumber int, partnerName string, quantity int) bool {
	if gameNumber < j.tolerance {
		return true
	}
	return j.midUpperPoint*gameNumber >= quantity*gameNumber
}
