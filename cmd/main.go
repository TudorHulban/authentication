package main

import (
	"context"

	"github.com/TudorHulban/authentication/fixtures"
	"github.com/TudorHulban/authentication/infra"
)

func main() {
	app := infra.InitializeApp(
		&configuration,
	)
	infra.InitializeTransportRoutes(app)

	ctx := context.Background()

	fixtures.InitializeAddTestUser(
		ctx,
		app,
	)

	app.Start()
}
