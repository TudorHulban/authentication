package srender

type ParamsCurrentMenuSidebar struct {
	TextLogo      string
	PathImageLogo string

	TextSection1 string

	TextSection1Entry1   string
	SymbolSection1Entry1 string
	URLSection1Entry1    string

	TextSection1Entry2   string
	SymbolSection1Entry2 string
	URLSection1Entry2    string
}

var ParamsMenuSidebarToUse = func(params *ParamsCurrentMenuSidebar) ParamsMenuSidebar {
	return ParamsMenuSidebar{
		TextLogo:      params.TextLogo,
		PathImageLogo: params.PathImageLogo,

		Sections: []*MenuSidebarSection{
			{
				TextSection: params.TextSection1,

				Entries: []*MenuSidebarSectionEntry{
					{
						TextSectionEntry: params.TextSection1Entry1,
						SymbolEntry:      params.SymbolSection1Entry1,
						URLEntry:         params.URLSection1Entry1,
					},
					{
						TextSectionEntry: params.TextSection1Entry2,
						SymbolEntry:      params.SymbolSection1Entry2,
						URLEntry:         params.URLSection1Entry2,
					},
				},
			},
		},
	}
}
