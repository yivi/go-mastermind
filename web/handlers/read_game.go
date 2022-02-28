package handlers

import (
	"githug.com/yivi/go-mastermind/lib"
	"net/http"
)

func ReadGame(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	game, ok := ctx.Value("game").(*lib.Game)
	if !ok {
		http.Error(w, http.StatusText(422), 422)
		return
	}

	guessData := &GuessData{Errors: make(map[string]string), Game: game}

	//goland:noinspection GoUnhandledErrorResult
	lib.Cn.GetTemplates()["index"].Execute(w, guessData)
}
