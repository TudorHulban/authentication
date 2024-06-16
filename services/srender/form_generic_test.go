package srender

import (
	"os"
	"testing"
)

func TestNewFormSearchTickets(t *testing.T) {
	p := ParamsNewFormSearchTickets{
		ActionForm:        "/tickets",
		LabelButtonSubmit: "Submit",
	}

	NewFormSearchTickets(&p).Render(os.Stdout)
}
