package pages

import (
	g "github.com/maragudk/gomponents"
	html "github.com/maragudk/gomponents/html"
)

var ScriptHTMX = html.Script(
	html.Src("https://unpkg.com/htmx.org@1.9.12"),
	g.Attr(
		"integrity", "sha384-ujb1lZYygJmzgSwoxRggbCHcjc0rB2XoQrxeTUQyRjrOnlCoYta87iKBWq3EsdM2",
	),
	g.Attr(
		"crossorigin",
		"anonymous",
	),
)

var LinkCSSWater = html.Link(
	html.Rel("stylesheet"),
	html.Href("https://cdn.jsdelivr.net/npm/water.css@2/out/dark.min.css"),
)
