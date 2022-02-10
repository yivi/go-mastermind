package main

import (
	"fmt"
	"githug.com/yivi/go-mastermind/lib"
	"regexp"
	"strings"
)

func main() {

	fmt.Println("Welcome to MasterMind Go!")
	fmt.Println("Match your wits against the unconquerable computer.")
	fmt.Println("")

	number := lib.Generate()

	for {
		fmt.Print("Choose a 4 digit number:")
		guess := readNumber()
		if match, _ := regexp.Match(`^\d{4}$`, []byte(guess)); !match {
			fmt.Println("💩 That doesn't look like a 4 digit number")
			continue
		}

		var guessArray [4]byte
		copy(guessArray[:], guess)
		good, regular := lib.CheckGuess(number, guessArray)

		g := strings.Repeat("🟩", int(good))
		r := strings.Repeat("🟨", int(regular))
		b := strings.Repeat("🟥", 4-int(good)-int(regular))

		fmt.Printf("%s%s%s: %d Good, %d Regular\n", g, r, b, good, regular)

		if good == 4 {
			fmt.Println("You WON!!!! 🎉🎉🎉")
			break
		}
	}
}

func readNumber() (guess string) {

	_, err := fmt.Scanln(&guess)
	if err != nil {
		guess = ""
	}

	return
}
