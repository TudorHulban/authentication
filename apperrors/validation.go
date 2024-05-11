package apperrors

import "fmt"

type ErrValidation struct {
	Issue  error
	Caller string
}

const areaErrServiceValidation = "Validation"

func (e ErrValidation) Error() string {
	var res [3]string

	res[0] = fmt.Sprintf("Area: %s", areaErrServiceValidation)
	res[1] = fmt.Sprintf("Caller: %s", e.Caller)

	res[2] = ""
	if e.Issue != nil {
		res[2] = fmt.Sprintf("Issue: %s", e.Issue.Error())
	}

	return res[0] + _space + res[1] + _space + res[2]
}

type ErrNilInput struct {
	InputName string
}

func (e ErrNilInput) Error() string {
	return fmt.Sprintf(
		"nil Input, name: %s",
		e.InputName,
	)
}

type ErrInvalidInput struct {
	InputName string
}

func (e ErrInvalidInput) Error() string {
	return fmt.Sprintf(
		"invalid Input, name: %s",
		e.InputName,
	)
}
