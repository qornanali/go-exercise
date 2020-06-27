package errorhandling

// InputError is ..
type InputError struct {
	Message string
}

func (e InputError) Error() string {
	return e.Message
}
