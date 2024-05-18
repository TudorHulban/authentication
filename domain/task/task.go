package task

import "github.com/TudorHulban/authentication/helpers"

type TaskMetadata struct {
	TimestampOfLastUpdate int64
	Status                TaskStatus
}

type TaskInfo struct {
	Name string

	*TaskMetadata

	OpenedByUserID uint
	Kind           uint8
}

type Task struct {
	helpers.PrimaryKey

	*TaskInfo
}
