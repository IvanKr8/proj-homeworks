package main

import (
	"context"
	"fmt"
	"time"
)

func SendToPlayers(roundCh chan Round, playerCh chan Player, ctx context.Context) {
	fmt.Println("Game `Kahoot` is start")
	fmt.Println()
	round := <-roundCh
	fmt.Println()
	fmt.Println("Received new round:")
	fmt.Println()
	fmt.Println("Question:", round.Question)
	fmt.Printf("Answers: ")
	for _, val := range round.Answers {
		fmt.Printf("%v ", val)
	}
	fmt.Println()
	time.Sleep(10 * time.Second)
	answeredPlayers := make(map[int]bool)
	for {
		select {
		case round = <-roundCh:
			for i := 1; i <= 5; i++ {
				if _, ok := answeredPlayers[i]; !ok {
					player := Player{
						Id:     i,
						Answer: round.Answers[i%len(round.Answers)],
					}
					playerCh <- player
					answeredPlayers[i] = true
				}
			}
		}
	}
}

func CountAnswers(playerCh chan Player, answerCountCh chan map[string]int, ctx context.Context, contextCancelFunc context.CancelFunc) {
	answerCount := make(map[string]int)
	answeredPlayers := make(map[int]bool)
	tempCount := make(map[string]int)

	for {
		select {
		case <-ctx.Done():
			return
		case player := <-playerCh:
			if _, ok := answeredPlayers[player.Id]; !ok {
				tempCount[player.Answer]++
				fmt.Printf("Player %d answered %s\n", player.Id, player.Answer)
				time.Sleep(100 * time.Millisecond)
				answeredPlayers[player.Id] = true
			}
			if len(answeredPlayers) == 5 {
				for answer, count := range tempCount {
					answerCount[answer] += count
				}
				fmt.Printf("-> %v", answerCount)
				contextCancelFunc()
				answerCountCh <- answerCount
				return
			}
		}
	}
}
