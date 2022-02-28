package handlers

import (
	"githug.com/yivi/go-mastermind/lib"
	"net/http"
)

func Index(w http.ResponseWriter, _ *http.Request) {

	templates := lib.Cn.GetTemplates()

	//goland:noinspection GoUnhandledErrorResult
	templates["index"].Execute(w, nil)
}
