package helpers

import (
	"os"
	"testing"

	"github.com/CloudyKit/jet"
	"github.com/stretchr/testify/require"
)

func TestUnixNanoTo(t *testing.T) {
	s := jet.NewHTMLSet(".")

	s.SetDevelopmentMode(true)

	templ, errGet := s.GetTemplate("timestamp_unixnanoto.jet")
	require.NoError(t, errGet)

	jetMap := jet.VarMap{}

	jetMap.Set(
		"UnixNanoTo",
		UnixNanoTo,
	)

	templ.Execute(
		os.Stdout,
		jetMap,
		nil,
	)
}
