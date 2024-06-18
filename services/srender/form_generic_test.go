package srender

import (
	"os"
	"testing"
)

func TestNewFormSearchTickets(t *testing.T) {
	p := ParamsNewFormSearchTickets{
		ActionButtonSearch: "/tickets",
		ActionButtonCreate: "/ticket",

		LabelButtonCreate: "Submit",
	}

	NewFormSearchTickets(&p).Render(os.Stdout)
}
