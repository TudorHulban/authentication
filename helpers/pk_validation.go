package helpers

import (
	"fmt"
	"strconv"

	"github.com/TudorHulban/authentication/apperrors"
)

type ParamsValidationStringPK struct {
	Value string

	Name   string
	Caller string
}

func ValidationStringID(params *ParamsValidationStringPK, result *uint) error {
	if len(params.Value) == 0 {
		return apperrors.ErrValidation{
			Caller: params.Caller,

			Issue: apperrors.ErrNilInput{
				InputName: params.Name,
			},
		}
	}

	numericValue, errConvToNumeric := strconv.Atoi(params.Value)
	if errConvToNumeric != nil {
		return apperrors.ErrValidation{
			Caller: params.Caller,

			Issue: fmt.Errorf(
				"%s: %w",
				params.Name,
				errConvToNumeric,
			),
		}
	}

	if numericValue < 0 {
		return apperrors.ErrValidation{
			Caller: params.Caller,

			Issue: apperrors.ErrNegativeInput{
				InputName: params.Name,
			},
		}
	}

	if numericValue == 0 {
		return apperrors.ErrValidation{
			Caller: params.Caller,

			Issue: apperrors.ErrZeroInput{
				InputName: params.Name,
			},
		}
	}

	*result = uint(numericValue)

	return nil
}
