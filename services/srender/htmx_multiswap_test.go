package srender

import (
	"fmt"
	"testing"

	g "github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

// TODO: proper test
func TestRenderNodes(t *testing.T) {
	n1 := html.P(
		g.Text("p1"),
	)

	n2 := html.P(
		g.Text("p2"),
	)

	fmt.Println(
		string(
			RenderNodes(
				n1,
				n2,
			),
		),
	)
}
