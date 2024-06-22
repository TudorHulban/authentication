package srender

import (
	g "github.com/maragudk/gomponents"
	html "github.com/maragudk/gomponents/html"
)

type MenuSidebarSectionEntry struct {
	TextSectionEntry string
	URLEntry         string

	HTMXMultiswapTargets []string

	SymbolEntry string
}

func (entry MenuSidebarSectionEntry) Render() g.Node {
	return html.Li(
		html.A(
			g.Attr(
				"href",
				entry.URLEntry,
			),

			g.Attr(
				"hx-ext",
				"multi-swap",
			),

			g.Attr(
				"hx-get",
				entry.URLEntry,
			),

			g.Attr(
				"hx-swap",
				newMultiswap(entry.HTMXMultiswapTargets),
			),

			html.Span(
				g.Attr(
					"class",
					"material-symbols-outlined",
				),

				g.Text(
					entry.SymbolEntry,
				),
			),

			g.Text(
				entry.TextSectionEntry,
			),
		),
	)
}

type MenuSidebarSection struct {
	TextSection string

	Entries []*MenuSidebarSectionEntry
}

func (section MenuSidebarSection) Render() []g.Node {
	result := []g.Node{
		html.H4(
			html.Span(
				g.Text(
					section.TextSection,
				),
			),

			html.Div(
				g.Attr(
					"class",
					"menu-separator",
				),
			),
		),
	}

	for _, entry := range section.Entries {
		result = append(result,
			entry.Render(),
		)
	}

	return result
}

// TODO: move to []g.Node
type MenuSidebar struct {
	PathImageLogo string
	TextLogo      string

	Sections []*MenuSidebarSection
}

type ParamsMenuSidebar struct {
	PathImageLogo string
	TextLogo      string

	Sections []*MenuSidebarSection
}

func NewMenuSidebar(params ParamsMenuSidebar) (*MenuSidebar, error) {
	// TODO inputvalidation

	return &MenuSidebar{
			PathImageLogo: params.PathImageLogo,
			TextLogo:      params.TextLogo,
			Sections:      params.Sections,
		},
		nil
}

func (menu MenuSidebar) Render() g.Node {
	sections := []g.Node{
		g.Attr(
			"class",
			"sidebar-links",
		),
	}

	for _, section := range menu.Sections {
		sections = append(sections,
			section.Render()...,
		)
	}

	return html.Aside(
		g.Attr(
			"class",
			"sidebar",
		),

		html.Div(
			g.Attr(
				"class",
				"sidebar-header",
			),

			html.Img(
				g.Attr(
					"alt",
					"logo",
				),

				g.Attr(
					"src",
					menu.PathImageLogo,
				),
			),

			html.H2(
				g.Text(menu.TextLogo),
			),
		),

		html.Ul(
			sections...,
		),
	)
}
