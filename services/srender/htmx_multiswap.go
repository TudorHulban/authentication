package srender

import (
	"bytes"
	"strings"

	"github.com/TudorHulban/authentication/helpers"
	g "github.com/maragudk/gomponents"
)

// no input validation.
func NewMultiswap(idsElements []string) string {
	idsSanitized := make([]string, len(idsElements), len(idsElements))

	for ix, id := range idsElements {
		idsSanitized[ix] = helpers.SanitizeCSSId(
			id,
		)
	}

	return strings.Join(idsSanitized, ",")
}

func RenderNodes(nodes ...g.Node) []byte {
	var buf bytes.Buffer

	for _, node := range nodes {
		node.Render(&buf)

		buf.WriteString("|")
	}

	return buf.Bytes()
}
