package types

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
