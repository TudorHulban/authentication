package main

import (
	"context"

	"github.com/TudorHulban/authentication/fixtures"
	"github.com/TudorHulban/authentication/infra"
)

func main() {
	app := infra.InitializeApp()

	ctx := context.Background()

	fixtures.InitializeAddTestUser(
		ctx,
		app,
	)
}
