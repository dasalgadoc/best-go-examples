package main

type StorierActor struct {
	BaseActor
	startReviewThreshold int
}

func NewStorierActor(role string, gameUpperBound int) *StorierActor {
	return &StorierActor{
		BaseActor: BaseActor{
			name:           "Storier",
			role:           role,
			gameUpperBound: gameUpperBound,
			quantityEarned: 0,
			gamesPlayed:    []Game{},
		},
		startReviewThreshold: 10,
	}
}

func (s *StorierActor) Propose(gameNumber int) (quantity int) {
	if gameNumber < s.startReviewThreshold {
		return s.gameUpperBound / 2
	}
	return int(float32(s.gameUpperBound) * s.calculateMultiplier(gameNumber))
}

func (s *StorierActor) Accept(gameNumber int, partnerName string, quantity int) bool {
	return quantity > int(float32(s.gameUpperBound)*s.calculateMultiplier(gameNumber))
}

func (s *StorierActor) calculateMultiplier(gameNumber int) float32 {
	numbersOfAccepted := 0
	for i := 0; i < gameNumber; i++ {
		if s.gamesPlayed[i].Accepted {
			numbersOfAccepted++
		}
	}

	if 1-float32(numbersOfAccepted)/float32(gameNumber) > 0.75 {
		return 0.75
	}

	return 1 - float32(numbersOfAccepted)/float32(gameNumber)
}
