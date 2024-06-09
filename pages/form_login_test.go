package pages

import (
	"os"
	"testing"
)

func TestFormLogin(t *testing.T) {
	form := FormLogin()

	form.Render(os.Stdout)
}
