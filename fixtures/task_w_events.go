package fixtures

import (
	"github.com/TudorHulban/authentication/helpers"
	"github.com/TudorHulban/authentication/services/stask"
)

type PiersFixtureTaskWEvents struct {
	ServiceTask *stask.Service
}

type ParamsFixtureTaskWEvents struct {
	TaskName     string
	NumberEvents uint
}

// TODO:
func FixtureTaskWEvents(piers *PiersFixtureTaskWEvents, params *ParamsFixtureTaskWEvents) helpers.PrimaryKey {
	return helpers.PrimaryKeyZero
}
