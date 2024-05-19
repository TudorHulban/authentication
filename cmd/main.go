package main

import (
	"context"
	"fmt"

	"github.com/TudorHulban/authentication/app"
	"github.com/TudorHulban/authentication/fixtures"
)

func main() {
	app := app.InitializeApp(
		&configuration,
	)

	ctx := context.Background()

	fixtures.FixtureAddTestUser(
		ctx,
		&fixtures.PiersFixtureAddTestUser{
			ServiceUser: app.ServiceUser,
		},
	)

	fmt.Println(
		app.Start(),
	)
}
