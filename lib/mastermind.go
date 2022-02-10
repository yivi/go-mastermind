package go_mastermind

import (
	"math/rand"
	"time"
)

func Generate() (numberArray [4]byte) {
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

	copy(numberArray[:], numbers[start:end])

	return

}

// CheckGuess Checks how many matches there are within `real` and `guess``
//
// Returns how many `goodGuesses` there were (right digit, right position), and how many "regular" guesses there were:
// (right digit, wrong position). When real=guess, goodGuesses -> 4 regularGuesses -> 0
//
func CheckGuess(real, guess [4]byte) (goodGuesses uint, regularGuesses uint) {

	regularGuesses = 0
	goodGuesses = 0

	for i := 0; i < 4; i++ {
		if real[i] == guess[i] {
			goodGuesses++
			continue
		} else {
			for j := 0; j < 4; j++ {
				if j == i {
					continue
				}
				if guess[i] == real[j] {
					regularGuesses++
					continue
				}
			}
		}
	}

	return

}
