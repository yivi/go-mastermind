package lib

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/spf13/viper"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Container struct {
	db             *sqlx.DB
	gameRepository *GameRepository
}

func (container *Container) GetConnection() *sqlx.DB {

	viper.SetDefault("db_host", "localhost")
	viper.SetDefault("db_port", 12790)
	viper.SetDefault("db_user", "superuser")
	viper.SetDefault("db_pass", "fake-password")
	viper.SetDefault("db_name", "mastermind")

	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		confFileError := viper.ReadInConfig() // Find and read the config file
		if confFileError != nil {             // Handle errors reading the config file
			panic(fmt.Errorf("Fatal error config file: %w \n", err))
		}
	}

	if container.db != nil {
		return container.db
	}

	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		viper.Get("db_host"),
		viper.Get("db_port"),
		viper.Get("db_user"),
		viper.Get("db_pass"),
		viper.Get("db_name"),
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
