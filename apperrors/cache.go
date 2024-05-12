package apperrors

const msgErrEntryNotFound = "entry not found"

type ErrEntryNotFound struct{}

func (ErrEntryNotFound) Error() string {
	return msgErrEntryNotFound
}

func (e ErrEntryNotFound) Is(err error) bool {
	_, couldCast := err.(ErrEntryNotFound)

	return couldCast
}
