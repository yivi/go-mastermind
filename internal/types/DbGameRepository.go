package types

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

const DateFormat = "2006-01-02 15:04:05"

type DbGameRepository struct {
	db *sqlx.DB
}

func (r *DbGameRepository) GetGameById(game *Game, Id string) error {
	sql := "SELECT * FROM games WHERE id = $1"
	statement, err := r.db.Preparex(sql)
	if err != nil {
		return &RepositoryError{message: fmt.Sprintf("Failed preparing statement: '%s'", sql), code: -1}
	}

	if err = statement.Get(game, Id); err != nil {
		return &RepositoryError{message: fmt.Sprintf("Game not found: '%s'", Id), code: 404}
	}

	return nil

}

func (r *DbGameRepository) AddGame(game *Game) error {

	var sql string
	getGameErr := r.GetGameById(&Game{}, game.Id)

	if getGameErr != nil {
		sql = "INSERT INTO games (id, number, created_at, updated_at, won, finished, guesses, guess_count) VALUES (:id, :number, :createdAt, :updatedAt, false, false, :guesses, :guessCount)"
	} else {
		sql = "UPDATE games set updated_at = :updatedAt, guesses = :guesses, guess_count = :guessCount, won = :won, finished = :finished WHERE id = :id"
	}

	stmt, prepareStmtErr := r.db.PrepareNamed(sql)
	if prepareStmtErr != nil {
		return &RepositoryError{message: "Failed preparing SQL statement: " + sql}
	}

	arg := map[string]interface{}{
		"id":         game.Id,
		"number":     game.Number,
		"guesses":    game.Guesses,
		"createdAt":  time.Now().Format(DateFormat),
		"updatedAt":  time.Now().Format(DateFormat),
		"guessCount": game.GuessCount,
		"won":        game.Won,
		"finished":   game.Finished,
	}
	_, getGameErr = stmt.Exec(arg)
	if getGameErr != nil {
		return &RepositoryError{message: fmt.Sprintf("Could not insert/update entity %s: %s", game.Id, getGameErr.Error())}
	}

	return nil
}

func (r *DbGameRepository) DeleteGame(ID string) error {
	sql := "DELETE FROM games WHERE id = $1"

	stmt, err := r.db.Preparex(sql)
	if err != nil {
		return &RepositoryError{message: "Failed preparing SQL statement: " + sql}
	}

	_, err = stmt.Exec(ID)
	if err != nil {
		return &RepositoryError{message: "Error deleting: " + sql}
	}

	return nil
}

func NewDbGameRepository(db *sqlx.DB) *DbGameRepository {
	return &DbGameRepository{db: db}
}
