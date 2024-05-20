package apperrors

import "fmt"

type ErrValidation struct {
	Issue  error
	Caller string
}

const areaErrValidation = "Validation"

func (e ErrValidation) Error() string {
	var res [3]string

	res[0] = fmt.Sprintf("Area: %s", areaErrValidation)
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

type ErrNegativeInput struct {
	InputName string
}

func (e ErrNegativeInput) Error() string {
	return fmt.Sprintf(
		"negative Input Name: %s",
		e.InputName,
	)
}

type ErrZeroInput struct {
	InputName string
}

func (e ErrZeroInput) Error() string {
	return fmt.Sprintf(
		"zero Input Name: %s",
		e.InputName,
	)
}
