package types

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/yivi/go-mastermind/web/templates"
	"html/template"
)

const (
	templatesDir = "views"
	layoutsDir   = "views/partials"
	templateGlob = "/*.tmpl"
)

type Container struct {
	Config *Config

	db             *sqlx.DB
	gameRepository GameRepositoryInterface

	templates map[string]*template.Template
}

func (c *Container) GetDatabase() *sqlx.DB {

	if c.Config == nil {
		panic("Not configured")
	}

	if c.db != nil {
		return c.db
	}

	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.Config.DbHost,
		c.Config.DbPort,
		c.Config.DbUser,
		c.Config.DbPass,
		c.Config.DbName,
	)

	db, openErr := sqlx.Open("postgres", psqlInfo)
	if openErr != nil {
		panic(openErr)
	}

	openErr = db.Ping()
	if openErr != nil {
		panic(openErr)
	}

	c.db = db
	c.db.MapperFunc(strcase.ToSnake)

	return c.db

}

func (c *Container) GetGameRepository() GameRepositoryInterface {
	if c.gameRepository != nil {
		return c.gameRepository
	}

	c.gameRepository = NewDbGameRepository(c.GetDatabase())

	return c.gameRepository

}

func (c *Container) GetTemplates() map[string]*template.Template {
	return c.templates
}

func (c *Container) LoadTemplates() error {
	if c.templates != nil {
		return nil
	}

	c.templates = make(map[string]*template.Template)

	var parseErr error
	c.templates["index"], parseErr = template.ParseFS(templates.ViewsFS, templatesDir+"/base.html.tmpl", templatesDir+"/index.html.tmpl", layoutsDir+templateGlob)
	c.templates["results"], parseErr = template.ParseFS(templates.ViewsFS, layoutsDir+"/results.html.tmpl")

	//c.templates["index"] = template.Must(template.ParseFiles(templatesDir+"/base.html.tmpl", templatesDir+"/index.html.tmpl"))
	//template.Must(c.templates["index"].ParseGlob(layoutsDir + templateGlob))
	//c.templates["results"] = template.Must(template.ParseFiles(layoutsDir + "/results.html.tmpl"))

	return parseErr
}

func NewContainer(conf *Config) *Container {
	return &Container{Config: conf}
}
