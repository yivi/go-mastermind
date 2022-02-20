package lib

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

const DateFormat = "2006-01-02 15:04:05"

type RepositoryError struct {
	message string
	code    int
}

func (e *RepositoryError) Error() string {
	return e.message
}

func (e *RepositoryError) GetCode() int {
	return e.code
}

type GameRepository struct {
	db *sqlx.DB
}

func (r *GameRepository) GetGameById(game *Game, ID string) error {
	sql := "SELECT * FROM games WHERE id = $1"
	statement, err := r.db.Preparex(sql)
	if err != nil {
		return &RepositoryError{message: fmt.Sprintf("Failed preparing statement: '%s'", sql), code: -1}
	}

	if err = statement.Get(game, ID); err != nil {
		return &RepositoryError{message: fmt.Sprintf("Game not found: '%s'", ID), code: 404}
	}

	return nil

}

func (r *GameRepository) AddGame(game *Game) error {

	var sql string
	getGameErr := r.GetGameById(&Game{}, game.Id)

	if getGameErr != nil {
		sql = "INSERT INTO games (id, number, created_at, updated_at, won, finished) VALUES (:id, :number, :createdAt, :updatedAt, false, false)"
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

func (r *GameRepository) DeleteGame(ID string) error {
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
