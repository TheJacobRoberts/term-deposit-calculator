package calculator

import (
	"fmt"
)

var (
	errValidationError = "failed to validate %s field: %s"
)

// ValidationError defines a custom error which happens during validation of
// user inputs. Intended to improve clarity and distinguish between 'actual'
// errors.
type ValidationError struct {
	err   error
	field string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf(errValidationError, e.field, e.err)
}
