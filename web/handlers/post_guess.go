package handlers

import (
	"githug.com/yivi/go-mastermind/lib"
	"net/http"
)

type GuessData struct {
	Game   *lib.Game
	Errors map[string]string
}

func PostGuess(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	game, ok := ctx.Value("game").(*lib.Game)
	if !ok {
		http.Error(w, http.StatusText(422), 422)
		return
	}
	guessData := &GuessData{Errors: make(map[string]string)}

	guess := lib.NewGuess(r.PostFormValue("guessNumber"))
	if !guess.Validate() {
		guessData.Errors["guess_invalid"] = "Invalid guess."
	} else {
		game.AddGuess(guess)
		addErr := lib.Cn.GetGameRepository().AddGame(game)
		if addErr != nil {
			guessData.Errors["save_err"] = "Could not save guess. Internal error."
		}
	}

	guessData.Game = game

	//goland:noinspection GoUnhandledErrorResult
	lib.Cn.GetTemplates()["index"].Execute(w, guessData)

}
