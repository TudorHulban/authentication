package app

import (
	"log"
	"os"

	"github.com/TudorHulban/authentication/apperrors"
	storememory "github.com/TudorHulban/authentication/infra/stores/store-memory"
	"github.com/TudorHulban/authentication/services/ssessions"
	"github.com/TudorHulban/authentication/services/suser"
)

func InitializeApp(config *ParamsNewApp) *App {
	app, errCr := NewApp(
		config,
		&PiersApp{
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
