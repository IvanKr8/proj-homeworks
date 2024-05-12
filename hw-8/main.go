package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
)

type Round struct {
	Question      string
	Answers       []string
	CorrectAnswer string
}

type Player struct {
	Id     int
	Answer string
}

func main() {
	roundCh := make(chan Round)
	roundReadyCh := make(chan Round)
	playerCh := make(chan Player)
	answerCountCh := make(chan map[string]int)

	ctx, contextCancelFunc := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer contextCancelFunc()

	go GenerateRound(roundCh, roundReadyCh, ctx)
	go SendToPlayers(roundCh, playerCh, ctx)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("All players are answered")
			fmt.Println("Round result:")
			answerCount := <-answerCountCh
			for answer, count := range answerCount {
				fmt.Printf("Answer %s was chosen by %d player(s)\n", answer, count)
			}
			return
		case <-roundReadyCh:
			go CountAnswers(playerCh, answerCountCh, ctx, contextCancelFunc)
		}
	}
}
