package lib

import (
	"strings"
	"time"
)

type GuessError struct {
}

func (ge *GuessError) Error() string {
	return "Invalid Guess"
}

type Game struct {
	Id         string
	Number     string
	Guesses    string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Won        bool
	Finished   bool
	GuessCount int
}

func (g *Game) GetGuesses() []string {
	return strings.Split(g.Guesses, "|")
}

func (g *Game) AddGuess(guess string) error {
	if Validate(guess) == false {
		return &GuessError{}
	}

	if g.Guesses != "" {
		g.Guesses += "|"
	}

	g.Guesses += guess
	g.GuessCount++

	return nil
}
