package types

type GuessError struct {
}

func (ge *GuessError) Error() string {
	return "Invalid Guess"
}
