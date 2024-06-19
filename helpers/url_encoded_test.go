package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorsProcessFormURLEncoded(t *testing.T) {
	testCases := []struct {
		description string
		input       string
		want        map[string]string
	}{
		{"1. error", "=", map[string]string{}},
		{"2. error", "&", map[string]string{}},
		{"3. error", "=&", map[string]string{}},
		{"4. error", "&=", map[string]string{}},
		{"5. error invalid field", "aa=&", map[string]string{}},
		{"6. error invalid field", "=aa&", map[string]string{}},
		{"7. one value", "aa=bb&", map[string]string{"aa": "bb"}},
		{"8. one value", "aa=bb", map[string]string{"aa": "bb"}},
		{"9. empty values", "id=&status=&name=", map[string]string{}},
		{"10. several values", "id=&status=&aa=bb", map[string]string{"aa": "bb"}},
	}

	for _, tc := range testCases {
		t.Run(tc.description,
			func(t *testing.T) {
				assert.Equal(t,
					tc.want,
					ProcessFormURLEncoded([]byte(tc.input)),
				)
			},
		)
	}
}
