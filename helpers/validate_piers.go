package helpers

import (
	"reflect"

	"github.com/TudorHulban/authentication/apperrors"
)

func ValidatePiers(piers any) error {
	if piers == nil {
		return apperrors.ErrValidation{
			Caller: "ValidatePiers",
			Issue: apperrors.ErrNilInput{
				InputName: "piers",
			},
		}
	}

	piersType := reflect.TypeOf(piers)
	piersValue := reflect.ValueOf(piers)

	piersKind := piersType.Kind()

	if piersKind == reflect.Ptr {
		piersType = piersType.Elem()
		piersValue = piersValue.Elem()
		piersKind = piersType.Kind()
	}

	switch piersKind {
	case reflect.Struct:
		for fieldIndex := 0; fieldIndex < piersType.NumField(); fieldIndex++ {
			field := piersType.Field(fieldIndex)

			switch field.Type.Kind() {
			case reflect.Ptr:
				if piersValue.Field(fieldIndex).IsNil() {
					return apperrors.ErrValidation{
						Caller: "ValidatePiers",
						Issue: apperrors.ErrNilInput{
							InputName: field.Name,
						},
					}
				}

			default:
				continue
			}
		}

	default:
		return apperrors.ErrValidation{
			Caller: "ValidatePiers",
			Issue: apperrors.ErrInvalidInput{
				InputName: "piers",
			},
		}
	}

	return nil
}
