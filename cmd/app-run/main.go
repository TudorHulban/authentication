package main

import (
	"context"
	"fmt"
	"os"

	"github.com/TudorHulban/authentication/app"
	"github.com/TudorHulban/authentication/apperrors"
	"github.com/TudorHulban/authentication/fixtures"
)

func main() {
	app, errInitialize := app.InitializeApp(
		&configuration,
	)
	if errInitialize != nil {
		fmt.Println(errInitialize)

		os.Exit(
			apperrors.OSExitForApplicationIssues,
		)
	}

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
