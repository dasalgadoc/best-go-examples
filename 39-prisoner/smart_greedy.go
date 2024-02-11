package main

type SmartGreedyActor struct {
	BaseActor
	retrospect    int
	testIncrement float32
}

func NewSmartGreedyActor(role string, gameUpperBound int) *SmartGreedyActor {
	return &SmartGreedyActor{
		BaseActor: BaseActor{
			name:           "SmartGreedy",
			role:           role,
			gameUpperBound: gameUpperBound,
			quantityEarned: 0,
			gamesPlayed:    []Game{},
		},
		retrospect:    3,
		testIncrement: 0.0,
	}
}

func (s *SmartGreedyActor) Propose(gameNumber int) (quantity int) {
	if gameNumber <= s.retrospect {
		return s.gameUpperBound / 10
	}

	if s.countFalseInRetrospect(gameNumber) > s.retrospect/2 {
		s.incrementTestIncrement()
		return int(float32(s.gameUpperBound) * s.testIncrement)
	}

	s.decrementTestIncrement()
	return int(float32(s.gameUpperBound) * s.testIncrement)
}

func (s *SmartGreedyActor) Accept(gameNumber int, partnerName string, quantity int) bool {
	if gameNumber <= s.retrospect {
		return quantity > 2*s.gameUpperBound/3
	}

	if s.countFalseInRetrospect(gameNumber) > s.retrospect/2 {
		return quantity >= s.gameUpperBound/2
	}

	return quantity > 2*s.gameUpperBound/3
}

func (s *SmartGreedyActor) countFalseInRetrospect(gameNumber int) int {
	count := 0
	for i := 1; i < s.retrospect; i++ {
		if !s.gamesPlayed[gameNumber-i].Accepted {
			count++
		}
	}
	return count
}

func (s *SmartGreedyActor) incrementTestIncrement() {
	if s.testIncrement >= 0.75 {
		s.testIncrement = 0.75
		return
	}
	s.testIncrement += 0.1
}

func (s *SmartGreedyActor) decrementTestIncrement() {
	if s.testIncrement <= 0.0 {
		s.testIncrement = 0.0
		return
	}
	s.testIncrement -= 0.1
}
