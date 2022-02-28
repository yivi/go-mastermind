package lib

import (
	"regexp"
	"strconv"
	"strings"
)

type Guess struct {
	Number  string
	Good    int
	Regular int
}

// Validate Checks that the received string is a valid MasterMind™️ number
//
// The `guess` needs to be 4 unique digits long, and should not start with 0.
func (g *Guess) Validate() bool {
	if match, _ := regexp.MatchString(`^[1-9]\d{3}$`, g.Number); !match {
		return false
	}

	for pos, char := range g.Number {
		for j := pos + 1; j < 4; j++ {
			if string(char) == string(g.Number[j]) {
				return false
			}
		}
	}

	return true
}

func NewGuess(guess string) *Guess {
	return &Guess{Number: guess}
}

func NewGuessFrom(full string) *Guess {
	guessParts := strings.Split(full, ",")
	guess := NewGuess(guessParts[0])
	guess.Good, _ = strconv.Atoi(guessParts[1])
	guess.Regular, _ = strconv.Atoi(guessParts[2])

	return guess
}
