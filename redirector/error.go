package redirector

type ValidationError struct {
	Message string
}

func (e ValidationError) Error() string {
	return e.Message
}
