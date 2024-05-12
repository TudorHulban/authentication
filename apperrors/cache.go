package apperrors

import "fmt"

type ErrEntryNotFound struct {
	Key any
}

func (e ErrEntryNotFound) Error() string {
	return fmt.Sprintf(
		"entry for key '%v' not found",
		e.Key,
	)
}
