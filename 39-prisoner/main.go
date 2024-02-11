package main

import "fmt"

func main() {
	its := 100
	gameUpperBound := 100

	tournament := struct {
		Length          int
		Proposers       []Actor
		ProposersPoints []int
		Acceptors       []Actor
		AcceptorsPoints []int
	}{
		Length: 7,
		Proposers: []Actor{
			NewAveragerActor("proposer", gameUpperBound),
			NewConformistActor("proposer", gameUpperBound),
			NewGreedyActor("proposer", gameUpperBound),
			NewJustifierActor("proposer", gameUpperBound),
			NewRandomerActor("proposer", gameUpperBound),
			NewStorierActor("proposer", gameUpperBound),
			NewSmartGreedyActor("proposer", gameUpperBound),
		},
		ProposersPoints: []int{0, 0, 0, 0, 0, 0, 0},
		Acceptors: []Actor{
			NewAveragerActor("acceptor", gameUpperBound),
			NewConformistActor("acceptor", gameUpperBound),
			NewGreedyActor("acceptor", gameUpperBound),
			NewJustifierActor("acceptor", gameUpperBound),
			NewRandomerActor("acceptor", gameUpperBound),
			NewStorierActor("acceptor", gameUpperBound),
			NewSmartGreedyActor("acceptor", gameUpperBound),
		},
		AcceptorsPoints: []int{0, 0, 0, 0, 0, 0, 0},
	}

	for i := 0; i < tournament.Length; i++ {
		for j := 0; j < tournament.Length; j++ {
			pPoints, aPoints := game(its, tournament.Proposers[i], tournament.Acceptors[j])
			tournament.ProposersPoints[i] += pPoints
			tournament.AcceptorsPoints[j] += aPoints
			tournament.Proposers[i].Restart()
			tournament.Acceptors[j].Restart()
		}
	}

	fmt.Printf("Tourament results:\n")
	fmt.Printf("%+v\n", tournament.ProposersPoints)
	fmt.Printf("%+v\n", tournament.AcceptorsPoints)
}

func game(its int, proposer, acceptor Actor) (int, int) {
	for i := 0; i < its; i++ {
		q := proposer.Propose(i)
		a := acceptor.Accept(i, proposer.GetName(), q)
		proposer.RegisterGame(Game{i, acceptor.GetName(), q, a})
		acceptor.RegisterGame(Game{i, proposer.GetName(), q, a})
	}

	return proposer.GetOverall(), acceptor.GetOverall()
}
