package types

import (
	"fmt"
	"github.com/rs/xid"
	"math/rand"
	"strings"
	"time"
)

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

const MaxGuesses = 10

func (g *Game) GetFifoGuesses() []*Guess {

	if g.Guesses == "" {
		return nil
	}

	guessStrings := strings.Split(g.Guesses, "|")

	guesses := make([]*Guess, 0)

	for _, guessString := range guessStrings {
		guesses = append(guesses, NewGuessFrom(guessString))
	}

	return guesses
}

func (g *Game) GetLifoGuesses() []*Guess {

	if g.Guesses == "" {
		return nil
	}

	guessStrings := strings.Split(g.Guesses, "|")

	guesses := make([]*Guess, 0)

	for i := len(guessStrings) - 1; i >= 0; i-- {
		guesses = append(guesses, NewGuessFrom(guessStrings[i]))
	}

	return guesses
}

func (g *Game) AddGuess(guess *Guess) {

	if g.GuessCount == MaxGuesses {
		return
	}

	g.checkGuess(guess)
	if g.Guesses != "" {
		g.Guesses += "|"
	}

	g.Guesses += fmt.Sprintf("%s,%d,%d", guess.Number, guess.Good, guess.Regular)
	g.GuessCount++

	if guess.Good == 4 {
		g.Won = true
		g.Finished = true
	}

	if g.GuessCount == MaxGuesses {
		g.Finished = true
	}
}

// CheckGuess Checks how many matches there are within `real` and `guess``
//
// Returns how many `goodGuesses` there were (right digit, right position), and how many "regular" guesses there were:
// (right digit, wrong position). When real=guess, goodGuesses -> 4 regularGuesses -> 0
//
func (g *Game) checkGuess(guess *Guess) {

	regularGuesses := 0
	goodGuesses := 0

	for i := 0; i < 4; i++ {
		if g.Number[i] == guess.Number[i] {
			goodGuesses++
			continue
		}

		if guess.Number[i] == g.Number[0] || guess.Number[i] == g.Number[1] || guess.Number[i] == g.Number[2] || guess.Number[i] == g.Number[3] {
			regularGuesses++
		}
	}

	guess.Good = goodGuesses
	guess.Regular = regularGuesses

}

func (g *Game) generateNumber() {
	rand.Seed(time.Now().Unix())

	numbers := []byte("1234567890")

	rand.Shuffle(len(numbers), func(i, j int) {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	})

	start := 0
	if string(numbers[start]) == "0" {
		start = 1
	}
	end := start + 4

	g.Number = string(numbers[start:end])
}

func NewGame() *Game {
	g := &Game{Id: xid.New().String()}
	g.generateNumber()

	return g
}
