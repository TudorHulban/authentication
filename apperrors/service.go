package apperrors

import "fmt"

type ErrServiceValidationNoUpdates struct {
	Caller string
}

const areaErrServiceValidationNoUpdates = "Service-Validation-NoUpdates"

func (e ErrServiceValidationNoUpdates) Error() string {
	return fmt.Sprintf("Area: %s", areaErrServiceValidationNoUpdates) +
		_space +
		fmt.Sprintf("Caller: %s", e.Caller)
}

type ErrServiceValidation struct {
	Issue  error
	Caller string
}

const areaErrServiceValidation = "Service-Validation"

func (e ErrServiceValidation) Error() string {
	var res [3]string

	res[0] = fmt.Sprintf("Area: %s", areaErrServiceValidation)
	res[1] = fmt.Sprintf("Caller: %s", e.Caller)

	res[2] = ""
	if e.Issue != nil {
		res[2] = fmt.Sprintf("Issue: %s", e.Issue.Error())
	}

	return res[0] + _space + res[1] + _space + res[2]
}

// TODO: add service name
type ErrService struct {
	Issue  error
	Caller string
}

const areaErrService = "Service"

func (e ErrService) Error() string {
	var res [3]string

	res[0] = fmt.Sprintf("Area: %s", areaErrService)
	res[1] = fmt.Sprintf("Caller: %s", e.Caller)

	res[2] = ""
	if e.Issue != nil {
		res[2] = fmt.Sprintf("Issue: %s", e.Issue.Error())
	}

	return res[0] + _space + res[1] + _space + res[2]
}
