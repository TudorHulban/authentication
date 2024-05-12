package main

import (
	"context"

	"github.com/TudorHulban/authentication/app"
	"github.com/TudorHulban/authentication/fixtures"
)

func main() {
	app := app.InitializeApp(
		&configuration,
	)

	ctx := context.Background()

	fixtures.InitializeAddTestUser(
		ctx,
		app,
	)

	app.Start()
}
