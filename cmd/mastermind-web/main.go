package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/yivi/go-mastermind/internal"
	"github.com/yivi/go-mastermind/internal/types"
	"github.com/yivi/go-mastermind/internal/web/handlers"
	"github.com/yivi/go-mastermind/internal/web/middleware"
	"github.com/yivi/go-mastermind/web/public"
	"net/http"
)

func init() {
	internal.Cf = types.NewConfig(false)
	internal.Cn = types.NewContainer(internal.Cf)

	templateError := internal.Cn.LoadTemplates()
	if templateError != nil {
		panic("Could not load templates.")
	}

}

func main() {

	r := chi.NewRouter()
	r.Use(chiMiddleware.Logger)

	staticFS := http.FS(public.Public)
	fs := http.FileServer(staticFS)

	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	r.Get("/", handlers.IndexHandler)

	r.Route("/{gameId:[a-z0-9]+}", func(r chi.Router) {
		r.Use(middleware.GameContextMiddleware)
		r.Get("/", handlers.ReadGameHandler)
		r.Post("/", handlers.PostGuessHandler)
	})

	//goland:noinspection GoUnhandledErrorResult
	http.ListenAndServe(fmt.Sprintf(":%d", internal.Cf.WebPort), r)

}
