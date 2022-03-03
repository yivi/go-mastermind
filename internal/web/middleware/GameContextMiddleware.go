package middleware

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/yivi/go-mastermind/internal"
	"github.com/yivi/go-mastermind/internal/types"
	"net/http"
)

func GameContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		gameId := chi.URLParam(r, "gameId")
		var game *types.Game

		if gameId == "new" {
			game = types.NewGame()
		} else {
			gameRepository := internal.Cn.GetGameRepository()
			game = &types.Game{}
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
