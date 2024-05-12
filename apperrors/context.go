package apperrors

import "fmt"

type ErrContextValueNotFound struct {
	Value string
}

func (e ErrContextValueNotFound) Error() string {
	return fmt.Sprintf(
		"value %s not found in context",
		e.Value,
	)
}
