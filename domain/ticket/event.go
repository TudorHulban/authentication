package ticket

import "github.com/TudorHulban/authentication/helpers"

type EventInfo struct {
	Content        string
	TimestampOfAdd int64
	OpenedByUserID uint
}

type Event struct {
	helpers.PrimaryKey

	*EventInfo
}
