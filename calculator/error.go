package calculator

import (
	"fmt"
)

var (
	errValidationError = "failed to validate %s field: %s"
)

type ValidationError struct {
	err   error
	field string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf(errValidationError, e.field, e.err)
}
