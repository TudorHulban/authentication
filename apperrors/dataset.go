package apperrors

import "fmt"

type ErrNoEntriesFound struct {
	Caller string
}

const areaErrNoEntriesFound = "Datasets"

func (e ErrNoEntriesFound) Error() string {
	var res [2]string

	res[0] = fmt.Sprintf("Area: %s", areaErrNoEntriesFound)
	res[1] = fmt.Sprintf("Caller: %s", e.Caller)

	return res[0] + _space + res[1]
}
