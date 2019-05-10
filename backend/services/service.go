package services

type ValidationError struct {
	message string
}

func NewValidationError(message string) *ValidationError {
	return &ValidationError{message}
}

func (e ValidationError) Error() string {
	return e.message
}

type InternalServerError struct {
	message string
}

func NewInternalServerError(message string) *InternalServerError {
	return &InternalServerError{message}
}

func (e InternalServerError) Error() string {
	return e.message
}
