package main

import (
	"context"
)

func GenerateRound(roundCh chan Round, roundReadyCh chan Round, ctx context.Context) {
	for {
		question := "What is the capital of Ukraine?"
		answers := []string{"Kiev", "London", "Berlin", "Madrid"}
		correctAnswer := "Kiev"
		round := Round{
			Question:      question,
			Answers:       answers,
			CorrectAnswer: correctAnswer,
		}
		roundCh <- round
		roundReadyCh <- round
	}
}
