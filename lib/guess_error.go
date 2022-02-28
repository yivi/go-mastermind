package lib

type GuessError struct {
}

func (ge *GuessError) Error() string {
	return "Invalid Guess"
}
