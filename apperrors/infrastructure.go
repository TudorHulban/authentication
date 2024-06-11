package apperrors

import "fmt"

type ErrInfrastructure struct {
	Issue              error
	NameInfrastructure string
	Caller             string
}

const areaErrInfrastructure = "Infrastructure"

func (e ErrInfrastructure) Error() string {
	var res [4]string

	res[0] = fmt.Sprintf("Area: %s", areaErrInfrastructure)
	res[1] = fmt.Sprintf("Name: %s", e.NameInfrastructure)
	res[2] = fmt.Sprintf("Caller: %s", e.Caller)

	res[3] = ""
	if e.Issue != nil {
		res[3] = fmt.Sprintf("Issue: %s", e.Issue.Error())
	}

	return res[0] + _space + res[1] + _space + res[2] + _space + res[3]
}
