package errors

type ValidationError struct {
	Errors map[string]string
}

func (e *ValidationError) Error() string {
    return "validation failed"
}