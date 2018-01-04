package make

// MultiError is an error consisting of multiple errors.
type MultiError struct {
	Errors []error
}

func (e *MultiError) Error() string {
	text := "Multiple errors occured:\n"
	for i, err := range e.Errors {
		text += err.Error()
		if i+1 < len(e.Errors) {
			text += "\n"
		}
	}
	return text
}
