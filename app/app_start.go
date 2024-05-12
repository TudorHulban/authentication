package app

func (a *App) Start() error {
	a.Transport.Use(
		[]string{
			RouteLogged,
		},
		a.MwAuthorization(),
	)

	InitializeTransportRoutes(a)

	return a.Transport.Listen(":" + a.port)
}
