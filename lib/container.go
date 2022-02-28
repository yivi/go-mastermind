package lib

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"githug.com/yivi/go-mastermind/web/templates"
	"html/template"
)

const (
	templatesDir = "views"
	layoutsDir   = "views/partials"
	templateGlob = "/*.tmpl"
)

type Container struct {
	Config         *Config
	db             *sqlx.DB
	gameRepository *GameRepository
	templates      map[string]*template.Template
}

func (c *Container) GetConnection() *sqlx.DB {

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

func (c *Container) GetGameRepository() *GameRepository {
	if c.gameRepository != nil {
		return c.gameRepository
	}

	c.gameRepository = &GameRepository{db: c.GetConnection()}

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

	return parseErr
}
