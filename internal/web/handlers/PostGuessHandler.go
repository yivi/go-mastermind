package handlers

import (
	"fmt"
	"github.com/yivi/go-mastermind/internal"
	"github.com/yivi/go-mastermind/internal/types"
	"net/http"
)

type GuessData struct {
	Game   *types.Game
	Errors map[string]string
}

func PostGuessHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	game, ok := ctx.Value("game").(*types.Game)
	if !ok {
		http.Error(w, http.StatusText(422), 422)
		return
	}
	guessData := &GuessData{Errors: make(map[string]string)}

	guess := types.NewGuess(r.PostFormValue("guessNumber"))
	if !guess.Validate() {
		guessData.Errors["guess_invalid"] = "Invalid guess."
	} else {
		game.AddGuess(guess)
	}

	if game.GuessCount > 0 {
		addErr := internal.Cn.GetGameRepository().AddGame(game)
		if addErr != nil {
			guessData.Errors["save_err"] = "Could not save guess. Internal error."
		}
	} else {
		game = &types.Game{Id: "new"}
	}

	guessData.Game = game

	//goland:noinspection GoUnhandledErrorResult
	http.Redirect(w, r, fmt.Sprintf("/%s", game.Id), http.StatusSeeOther)

}
