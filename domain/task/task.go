package task

import "github.com/TudorHulban/authentication/helpers"

type TaskInfo struct {
	Name string

	TimestampOfLastUpdate int64
	OpenedByUserID        uint
	Kind                  uint8
	Status                TaskStatus
}

type Task struct {
	helpers.PrimaryKey
	TaskInfo
}
