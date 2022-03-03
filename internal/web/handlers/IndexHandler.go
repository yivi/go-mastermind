package handlers

import (
	"github.com/yivi/go-mastermind/internal"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, _ *http.Request) {

	templates := internal.Cn.GetTemplates()

	//goland:noinspection GoUnhandledErrorResult
	templates["index"].Execute(w, nil)
}
