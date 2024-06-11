package pages

import (
	"os"
	"testing"
)

func TestFormLogin(t *testing.T) {
	form := FormLogin(t.Name())

	form.Render(os.Stdout)
}
