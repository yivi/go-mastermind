package handlers

import (
	"github.com/yivi/go-mastermind/internal"
	"github.com/yivi/go-mastermind/internal/types"
	"net/http"
)

func ReadGameHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	game, ok := ctx.Value("game").(*types.Game)
	if !ok {
		http.Error(w, http.StatusText(422), 422)
		return
	}

	guessData := &GuessData{Errors: make(map[string]string), Game: game}

	//goland:noinspection GoUnhandledErrorResult
	internal.Cn.GetTemplates()["index"].Execute(w, guessData)
}
