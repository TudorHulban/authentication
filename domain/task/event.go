package task

import "github.com/TudorHulban/authentication/helpers"

type EventInfo struct {
	Content string

	OpenedByUserID uint
}

type Event struct {
	helpers.PrimaryKey
	EventInfo
}
