package srender

import (
	g "github.com/maragudk/gomponents"
	html "github.com/maragudk/gomponents/html"
)

// see https://favicon.io/favicon-converter/

var LinksFavicon = []g.Node{
	html.Link(
		html.Rel("apple-touch-icon"),
		html.Href("/public/apple-touch-icon.png"),

		g.Attr(
			"sizes",
			"180x180",
		),
	),

	html.Link(
		html.Rel("icon"),
		html.Href("/public/favicon-32x32.png"),

		g.Attr(
			"type",
			"image/png",
		),

		g.Attr(
			"sizes",
			"32x32",
		),
	),

	html.Link(
		html.Rel("icon"),
		html.Href("/public/favicon-16x16.png"),

		g.Attr(
			"type",
			"image/png",
		),

		g.Attr(
			"sizes",
			"16x16",
		),
	),

	html.Link(
		html.Rel("manifest"),
		html.Href("/public/site.webmanifest"),
	),
}
