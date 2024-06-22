package srender

import (
	g "github.com/maragudk/gomponents"
	html "github.com/maragudk/gomponents/html"
)

var LinkCSSWater = html.Link(
	html.Rel("stylesheet"),
	html.Href("https://cdn.jsdelivr.net/npm/water.css@2/out/dark.min.css"),
)

var LinkCSSCommon = html.Link(
	html.Rel("stylesheet"),
	html.Href("/public/common.css"),
)

var LinkCSSMaterialSymbolOutlined = html.Link(
	html.Rel("stylesheet"),
	html.Href("https://fonts.googleapis.com/css2?family=Material+Symbols+Outlined:opsz,wght,FILL,GRAD@20..48,100..700,0..1,-50..200"),
)

var hidden = g.Attr(
	"hidden",
)

// var ScriptHTMX = html.Script(
// 	html.Src("https://unpkg.com/htmx.org@2.0.0"),
// 	g.Attr(
// 		"integrity", "sha384-wS5l5IKJBvK6sPTKa2WZ1js3d947pvWXbPJ1OmWfEuxLgeHcEbjUUA5i9V5ZkpCw",
// 	),
// 	g.Attr(
// 		"crossorigin",
// 		"anonymous",
// 	),
// )

var ScriptHTMX = html.Script(
	html.Src(
		"/public/htmx.js",
	),
)

// var ScriptHTMXMultiswap = html.Script(
// 	html.Src(
// 		"https://unpkg.com/htmx-ext-multi-swap@2.0.0/multi-swap.js",
// 	),
// )

var ScriptHTMXMultiswap = html.Script(
	html.Src(
		"/public/multi-swap.js",
	),
)

var ScriptCommonJS = html.Script(
	html.Src("/public/common.js"),
)
