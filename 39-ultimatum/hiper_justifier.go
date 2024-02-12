package main

type JustifierWithOutToleranceActor struct {
	BaseActor
}

func NewJustifierWithOutToleranceActor(role string, gameUpperBound int) *JustifierWithOutToleranceActor {
	return &JustifierWithOutToleranceActor{
		BaseActor: BaseActor{
			name:           "HighJustifier",
			role:           role,
			gameUpperBound: gameUpperBound,
			quantityEarned: 0,
			gamesPlayed:    []Game{},
		},
	}
}

func (j *JustifierWithOutToleranceActor) Propose(_ int) (quantity int) {
	return j.gameUpperBound / 2
}

func (j *JustifierWithOutToleranceActor) Accept(_ int, _ string, quantity int) bool {
	if quantity >= j.gameUpperBound/2 {
		return true
	}
	return false
}
