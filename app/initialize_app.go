package app

import (
	"context"
	"log"
	"os"

	"github.com/TudorHulban/authentication/apperrors"
	"github.com/TudorHulban/authentication/domain/task"
	"github.com/TudorHulban/authentication/fixtures"
	storefilefixtures "github.com/TudorHulban/authentication/fixtures/store-file-fixtures"
	storefile "github.com/TudorHulban/authentication/infra/stores/store-file"
	storememory "github.com/TudorHulban/authentication/infra/stores/store-memory"
	"github.com/TudorHulban/authentication/services/ssessions"
	"github.com/TudorHulban/authentication/services/stask"
	"github.com/TudorHulban/authentication/services/suser"
)

func InitializeApp(config *ParamsNewApp) *App {
	app, errCr := NewApp(
		config,
		&PiersApp{
			ServiceUser: suser.NewService(
				storememory.NewStoreMemory(),
			),

			ServiceTask: stask.NewService(
				// storememory.NewStoreTask(),
				storefile.NewStoreTask(&storefilefixtures.ParamsStoreFile),
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

	fixtures.FixtureTaskWEvents(
		context.Background(),
		&fixtures.PiersFixtureTaskWEvents{
			ServiceTask: app.serviceTask,
		},
		&fixtures.ParamsFixtureTaskWEvents{
			TaskName:           "task 1",
			TaskKind:           task.KindUndefined,
			TaskOpenedByUserID: 1,
			NumberEvents:       10,
		},
	)

	return app
}
