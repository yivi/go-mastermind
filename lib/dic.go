package lib

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Container struct {
	Config         *Config
	db             *sqlx.DB
	gameRepository *GameRepository
}

func (container *Container) GetConnection() *sqlx.DB {

	if container.db != nil {
		return container.db
	}

	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		container.Config.DbHost,
		container.Config.DbPort,
		container.Config.DbUser,
		container.Config.DbPass,
		container.Config.DbName,
	)

	db, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	container.db = db
	container.db.MapperFunc(strcase.ToSnake)

	return container.db

}

func (container *Container) GetGameRepository() *GameRepository {
	if container.gameRepository != nil {
		return container.gameRepository
	}

	container.gameRepository = &GameRepository{db: container.GetConnection()}

	return container.gameRepository

}
