package app

func (a *App) root() string {
	return ":" + a.port
}

func (a *App) baseURL() string {
	return a.host + a.root()
}

func (a *App) Start() error {
	a.Transport.Use(
		[]string{
			RouteLogged,
			RouteTask,
			RouteTasks,
		},
		a.MwAuthentication(),
	)

	InitializeTransportRoutes(a)

	return a.Transport.Listen(a.root())
}
