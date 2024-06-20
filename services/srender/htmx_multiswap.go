package srender

import (
	"strings"

	"github.com/TudorHulban/authentication/helpers"
)

func newMultiswap(idsElements []string) string {
	idsSanitized := make([]string, len(idsElements), len(idsElements))

	for ix, id := range idsElements {
		idsSanitized[ix] = helpers.SanitizeCSSId(
			id,
		)
	}

	return "multi:" + strings.Join(idsSanitized, ",")
}
