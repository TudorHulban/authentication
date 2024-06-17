package helpers

import (
	"fmt"
	"testing"
)

// TODO proper test
func TestSanitizeCSSIndex(t *testing.T) {
	fmt.Println(
		SanitizeCSSId("##el"),
	)
	fmt.Println(
		SanitizeCSSId("el"),
	)
}
