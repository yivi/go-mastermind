package main

import (
	"embed"
	"fmt"
	"githug.com/yivi/go-mastermind/lib"
	"net/http"
	"text/template"
)

type IndexData struct {
}

var (
	// viewsFS is our viewsFS web server content
	//go:embed views
	viewsFS   embed.FS
	templates map[string]*template.Template
	config    lib.Config
	container lib.Container
)

const (
	templatesDir = "views"
	layoutsDir   = "views/partials"
	templateGlob = "/*.tmpl"
)

func loadTemplates() error {
	if templates != nil {
		return nil
	}

	templates = make(map[string]*template.Template)

	var parseErr error
	templates["index"], parseErr = template.ParseFS(viewsFS, templatesDir+"/base.html.tmpl", templatesDir+"/index.html.tmpl", layoutsDir+templateGlob)

	return parseErr
}

func init() {
	err := loadTemplates()
	if err != nil {
		panic("Could not load templates")
	}

	config.Init(false)
	container.Config = &config

}

func indexHandler(w http.ResponseWriter, c *http.Request) {

	//goland:noinspection GoUnhandledErrorResult
	templates["index"].Execute(w, &IndexData{})
}

func guessHandler(writer http.ResponseWriter, request *http.Request) {

}

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/guess", guessHandler)

	//goland:noinspection GoUnhandledErrorResult
	http.ListenAndServe(fmt.Sprintf(":%d", config.WebPort), nil)

}
