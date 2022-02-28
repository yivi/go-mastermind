package main

import (
	"fmt"
	"githug.com/yivi/go-mastermind/lib"
	"strings"
)

var game *lib.Game

func main() {

	lib.Cf.Initialize(false)
	lib.Cn.Config = &lib.Cf

	gameRepository := lib.Cn.GetGameRepository()

	fmt.Println("Welcome to MasterMind Go!")
	fmt.Println("Match your wits against the unconquerable computer.")
	fmt.Println("")

	for {
		fmt.Print("Want to pick-up a game in progress? Leave blank to start fresh: ")
		gameId := readInput()

		if gameId != "" {
			getGameErr := gameRepository.GetGameById(game, gameId)
			if getGameErr != nil {
				fmt.Println("Error: " + getGameErr.Error())
				continue
			}

			if game.Won || game.Finished {
				fmt.Println("That game is not playable.")
				continue
			}

			break
		}

		game = lib.NewGame()
		err := gameRepository.AddGame(game)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		break

	}

	fmt.Println("Cheating: " + game.Number)
	for {
		fmt.Print("Choose a 4 digit number:")
		guessNumber := readInput()
		guess := lib.NewGuess(guessNumber)

		if !guess.Validate() {
			fmt.Println("💩 That doesn't look like a  VALID 4 digit number")
			continue
		}

		game.AddGuess(guess)

		g := strings.Repeat("🟩", guess.Good)
		r := strings.Repeat("🟨", guess.Regular)
		b := strings.Repeat("🟥", 4-guess.Good-guess.Regular)

		fmt.Printf("%s%s%s: %d Good, %d Regular\n", g, r, b, guess.Good, guess.Regular)

		if guess.Good == 4 {
			fmt.Println("You WON!!!! 🎉🎉🎉")
			game.Won = true
			game.Finished = true
		}

		if game.GuessCount == 9 {
			fmt.Println("You took too long, you lose! 👎👎👎")
			game.Finished = true
		}

		addGameErr := gameRepository.AddGame(game)
		if addGameErr != nil {
			fmt.Println("Could not save game 😢: " + addGameErr.Error())
			continue
		}

		if game.Finished {
			break
		}
	}
}

func readInput() (guess string) {

	_, err := fmt.Scanln(&guess)
	if err != nil {
		guess = ""
	}

	return
}
