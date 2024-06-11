package helpers

import (
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {
	fmt.Println(
		NewWordFrom("xxx").
			Hash(
				NewWordFrom("1"),
			),
	)
}
