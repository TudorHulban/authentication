package apperrors

import "fmt"

type ErrNoEntriesFound struct {
	Caller string
}

const areaErrNoEntriesFound = "Datasets"
const msgErrNoEntriesFound = "no entries found"

func (e ErrNoEntriesFound) Error() string {
	var res [2]string

	res[0] = fmt.Sprintf("Area: %s", areaErrNoEntriesFound)
	res[1] = fmt.Sprintf("Caller: %s", e.Caller)

	return res[0] + _space + res[1] + ": " + msgErrNoEntriesFound
}

type ErrEntryAlreadyExists struct {
	Entry  any
	Caller string
}

const areaErrEntryAlreadyExists = "Datasets"

func (e ErrEntryAlreadyExists) Error() string {
	var res [3]string

	res[0] = fmt.Sprintf("Area: %s", areaErrEntryAlreadyExists)
	res[1] = fmt.Sprintf("Caller: %s", e.Caller)
	res[2] = fmt.Sprintf("Entry: %#v", e.Entry)

	return res[0] + _space + res[1] + _space + res[2]
}
