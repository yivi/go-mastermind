package handlers

import (
	"context"
	"github.com/go-chi/chi/v5"
	"githug.com/yivi/go-mastermind/lib"
	"net/http"
)

func GameMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		gameId := chi.URLParam(r, "gameId")
		var game *lib.Game

		if gameId == "new" {
			game = lib.NewGame()
		} else {
			gameRepository := lib.Cn.GetGameRepository()
			game = &lib.Game{}
			getErr := gameRepository.GetGameById(game, gameId)
			if getErr != nil {
				http.Error(w, http.StatusText(404), 404)
				return
			}
		}

		ctx := context.WithValue(r.Context(), "game", game)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
