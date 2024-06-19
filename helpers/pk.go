package helpers

import (
	"strconv"

	"github.com/TudorHulban/authentication/apperrors"
)

type PrimaryKey uint64

func NewPrimaryKey(value string) (PrimaryKey, error) {
	numericPK, errConv := strconv.ParseUint(value, 10, 64)
	if errConv != nil {
		return PrimaryKeyZero,
			apperrors.ErrServiceValidation{
				Issue:  errConv,
				Caller: "NewPrimaryKey",
			}
	}

	return PrimaryKey(
			numericPK,
		),
		nil
}

func (pk PrimaryKey) String() string {
	return strconv.FormatUint(
		uint64(pk),
		10,
	)
}

const PrimaryKeyZero = PrimaryKey(0)
