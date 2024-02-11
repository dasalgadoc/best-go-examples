package main

import "math/rand"

type RandomerActor struct {
	BaseActor
}

func NewRandomerActor(role string, gameUpperBound int) *RandomerActor {
	return &RandomerActor{
		BaseActor: BaseActor{
			name:           "Randomer",
			role:           role,
			gameUpperBound: gameUpperBound,
			quantityEarned: 0,
			gamesPlayed:    []Game{},
		},
	}
}

func (r *RandomerActor) Propose(gameNumber int) (quantity int) {
	return rand.Intn(r.gameUpperBound)
}

func (r *RandomerActor) Accept(gameNumber int, partnerName string, quantity int) bool {
	return rand.Intn(100) <= 50
}
