package helpers

import (
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {
	fmt.Println(
		Hash(
			[]byte("xxx"),
			NewWordFrom("1"),
		),
	)
}
