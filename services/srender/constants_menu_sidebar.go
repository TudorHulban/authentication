package srender

type ParamsCurrentMenuSidebar struct {
	TextLogo      string
	PathImageLogo string

	TextSection1 string

	TextSection1Entry1        string
	SymbolSection1Entry1      string
	URLSection1Entry1         string
	HTMXTargetsSection1Entry1 []string

	TextSection1Entry2        string
	SymbolSection1Entry2      string
	URLSection1Entry2         string
	HTMXTargetsSection1Entry2 []string

	TextSection2 string

	TextSection2Entry1   string
	SymbolSection2Entry1 string
	URLSection2Entry1    string

	TextSection2Entry2   string
	SymbolSection2Entry2 string
	URLSection2Entry2    string
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
						TextSectionEntry:     params.TextSection1Entry1,
						SymbolEntry:          params.SymbolSection1Entry1,
						URLEntry:             params.URLSection1Entry1,
						HTMXMultiswapTargets: params.HTMXTargetsSection1Entry1,
					},
					{
						TextSectionEntry:     params.TextSection1Entry2,
						SymbolEntry:          params.SymbolSection1Entry2,
						URLEntry:             params.URLSection1Entry2,
						HTMXMultiswapTargets: params.HTMXTargetsSection1Entry2,
					},
				},
			},

			{
				TextSection: params.TextSection2,

				Entries: []*MenuSidebarSectionEntry{
					{
						TextSectionEntry: params.TextSection2Entry1,
						SymbolEntry:      params.SymbolSection2Entry1,
						URLEntry:         params.URLSection2Entry1,
					},
					{
						TextSectionEntry: params.TextSection2Entry2,
						SymbolEntry:      params.SymbolSection2Entry2,
						URLEntry:         params.URLSection2Entry2,
					},
				},
			},
		},
	}
}
