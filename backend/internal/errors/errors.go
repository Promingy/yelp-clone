package errors

type ValidationError struct {
	Errors map[string]string
}

func (e *ValidationError) Error() string {
    return "validation failed"
}

/// Response Types
type SingleErrorResponse map[string]string
type MultiErrorResponse map[string]map[string]string

func NewSingleError(msg string) SingleErrorResponse {
	return SingleErrorResponse{"error": msg}
}

func NewMultiError(errors map[string]string) MultiErrorResponse {
	return MultiErrorResponse{"errors": errors}	
}