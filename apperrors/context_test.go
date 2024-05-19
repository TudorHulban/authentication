package apperrors

import (
	"fmt"
	"testing"
)

func TestErrContextValueNotFound(t *testing.T) {
	fmt.Println(
		ErrContextValueNotFound{
			Value: "xxx",
		},
	)
}
