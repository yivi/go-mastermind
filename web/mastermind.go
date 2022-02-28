package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"githug.com/yivi/go-mastermind/lib"
	"githug.com/yivi/go-mastermind/web/handlers"
	"net/http"
)

func init() {
	lib.Cf.Initialize(false)
	lib.Cn.Config = &lib.Cf
	templateError := lib.Cn.LoadTemplates()
	if templateError != nil {
		panic("Could not load templates.")
	}

}

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", handlers.Index)
	r.Route("/{gameId:[a-z0-9]+}", func(r chi.Router) {
		r.Use(handlers.GameMiddleware)
		r.Get("/", handlers.ReadGame)
		r.Post("/", handlers.PostGuess)
	})

	//goland:noinspection GoUnhandledErrorResult
	http.ListenAndServe(fmt.Sprintf(":%d", lib.Cf.WebPort), r)

}
