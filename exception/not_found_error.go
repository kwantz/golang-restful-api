package exception

type NotFoundError struct {
	Error string
}

func NewNotFoundError(err error) NotFoundError {
	return NotFoundError{Error: err.Error()}
}

func PanicNotFoundIfError(err error) {
	if err != nil {
		panic(NewNotFoundError(err))
	}
}
