package helpers

import (
	"fmt"
	"testing"
)

func TestUnixNanoTo(t *testing.T) {
	fmt.Println(
		UnixNanoTo(50000000000000),
	)
}
