package types

type GameRepositoryInterface interface {
	GetGameById(game *Game, Id string) error
	AddGame(game *Game) error
	DeleteGame(ID string) error
}
