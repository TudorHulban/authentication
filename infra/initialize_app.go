package infra

import (
	"log"
	"os"

	"github.com/TudorHulban/authentication/app"
	"github.com/TudorHulban/authentication/apperrors"
	storememory "github.com/TudorHulban/authentication/infra/stores/store-memory"
	"github.com/TudorHulban/authentication/services/ssessions"
	"github.com/TudorHulban/authentication/services/suser"
)

func InitializeApp(config *app.ParamsNewApp) *app.App {
	app, errCr := app.NewApp(
		config,
		&app.PiersApp{
			ServiceUser: suser.NewService(
				storememory.NewStoreMemory(),
			),

			ServiceSessions: ssessions.NewService(),
		},
	)
	if errCr != nil {
		log.Printf(
			"InitializeApp: %v",
			errCr,
		)

		os.Exit(
			apperrors.OSExitForApplicationIssues,
		)
	}

	return app
}
