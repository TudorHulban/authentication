package srender

import (
	"context"

	"github.com/TudorHulban/authentication/app/constants"
	g "github.com/maragudk/gomponents"
	html "github.com/maragudk/gomponents/html"
)

type ParamsTableItems struct {
	TableHead g.Node
	TableBody g.Node
}

func (s *Service) TableItems(ctx context.Context, params *ParamsTableItems) g.Node {
	return html.Div(
		html.Table(
			[]g.Node{
				g.Attr(
					"id",
					constants.IDItemsTable,
				),

				g.Attr(
					"class",
					constants.ClassTableItems,
				),

				params.TableHead,
				params.TableBody,
			}...,
		),
	)
}
