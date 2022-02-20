package lib

import (
	"math/rand"
	"regexp"
	"time"
)

func Generate() string {
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

	return string(numbers[start:end])

}

// CheckGuess Checks how many matches there are within `real` and `guess``
//
// Returns how many `goodGuesses` there were (right digit, right position), and how many "regular" guesses there were:
// (right digit, wrong position). When real=guess, goodGuesses -> 4 regularGuesses -> 0
//
func CheckGuess(real, guess string) (goodGuesses uint, regularGuesses uint) {

	regularGuesses = 0
	goodGuesses = 0

	for i := 0; i < 4; i++ {
		if real[i] == guess[i] {
			goodGuesses++
			continue
		}

		if guess[i] == real[0] || guess[i] == real[1] || guess[i] == real[2] || guess[i] == real[3] {
			regularGuesses++
		}
	}

	return
}

// Validate Checks that the received string is a valid MasterMind™️ number
//
// The `guess` needs to be 4 unique digits long, and should not start with 0.
func Validate(guess string) bool {
	if match, _ := regexp.MatchString(`^\d{4}$`, guess); !match {
		return false
	}
	if string(guess[0]) == "0" {
		return false
	}

	for pos, char := range guess {
		for j := pos + 1; j < 4; j++ {
			if string(char) == string(guess[j]) {
				return false
			}
		}
	}

	return true
}
