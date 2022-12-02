package part2

import (
	"bufio"
	"io"
	"strings"
)

/*
A = Rock
B = Paper
C = Scissors

X = Lose
Y = Draw
Z = Win

Total Score = sum of your scores for each round

Rock = 1
Paper = 2
Scissors = 3

Lost = 0
Draw = 3
Win = 6

*/

const (
	ROCK     = "ROCK"
	PAPER    = "PAPER"
	SCISSORS = "SCISSORS"
)

const (
	LOSE = "X"
	DRAW = "Y"
	WIN  = "Z"
)

const (
	ROCK_SCORE = iota + 1
	PAPER_SCORE
	SCISSORS_SCORE
)

const (
	LOST_SCORE = 0
	DRAW_SCORE = 3
	WIN_SCORE  = 6
)

var selection map[string]string = map[string]string{
	"A": ROCK,
	"B": PAPER,
	"C": SCISSORS,
}

var selectionScore map[string]int = map[string]int{
	ROCK:     ROCK_SCORE,
	PAPER:    PAPER_SCORE,
	SCISSORS: SCISSORS_SCORE,
}

func roundScore(oppChoice string, myChoice string) int {
	score := selectionScore[myChoice]
	if oppChoice == myChoice {
		score += DRAW_SCORE
		return score
	}
	switch oppChoice {
	case ROCK:
		if myChoice == PAPER {
			score += WIN_SCORE
		}
	case PAPER:
		if myChoice == SCISSORS {
			score += WIN_SCORE
		}
	case SCISSORS:
		if myChoice == ROCK {
			score += WIN_SCORE
		}
	}
	return score
}

func getMyChoice(oppChoice string, outcome string) string {
	if outcome == DRAW {
		return oppChoice
	}

	var myChoice string

	switch oppChoice {
	case ROCK:
		if outcome == WIN {
			myChoice = PAPER
		} else {
			myChoice = SCISSORS
		}
	case PAPER:
		if outcome == WIN {
			myChoice = SCISSORS
		} else {
			myChoice = ROCK
		}
	case SCISSORS:
		if outcome == WIN {
			myChoice = ROCK
		} else {
			myChoice = PAPER
		}
	}
	return myChoice
}

func Run(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	totalScore := 0
	for scanner.Scan() {
		if scanner.Err() != nil {
			panic(scanner.Err())
		}
		cols := strings.Fields(scanner.Text())
		oppChoice := selection[cols[0]]
		desiredOutcome := cols[1]
		totalScore += roundScore(oppChoice, getMyChoice(oppChoice, desiredOutcome))
	}
	return totalScore
}
