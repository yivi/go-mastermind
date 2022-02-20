package main

import (
	"fmt"
	"github.com/rs/xid"
	"githug.com/yivi/go-mastermind/lib"
	"strings"
)

func main() {

	container := lib.Container{}
	gameRepository := container.GetGameRepository()

	var game lib.Game

	fmt.Println("Welcome to MasterMind Go!")
	fmt.Println("Match your wits against the unconquerable computer.")
	fmt.Println("")

	for {
		fmt.Print("Want to pick-up a game in progress? Leave blank to start fresh: ")
		gameId := readInput()

		if gameId != "" {
			getGameErr := gameRepository.GetGameById(&game, gameId)
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

		game = lib.Game{
			Id:     xid.New().String(),
			Number: lib.Generate(),
		}
		err := gameRepository.AddGame(&game)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		break

	}

	fmt.Println("Cheating: " + game.Number)
	for {
		fmt.Print("Choose a 4 digit number:")
		guess := readInput()

		addGuessErr := game.AddGuess(guess)
		if addGuessErr != nil {
			fmt.Println("💩 That doesn't look like a  VALID 4 digit number")
			continue
		}

		good, regular := lib.CheckGuess(game.Number, guess)

		g := strings.Repeat("🟩", int(good))
		r := strings.Repeat("🟨", int(regular))
		b := strings.Repeat("🟥", 4-int(good)-int(regular))

		fmt.Printf("%s%s%s: %d Good, %d Regular\n", g, r, b, good, regular)

		if good == 4 {
			fmt.Println("You WON!!!! 🎉🎉🎉")
			game.Won = true
			game.Finished = true
		}

		if game.GuessCount == 9 {
			fmt.Println("You too too long, you lose! 👎👎👎")
			game.Finished = true
		}

		addGameErr := gameRepository.AddGame(&game)
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
