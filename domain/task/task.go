package task

type TaskMetadata struct {
	TimestampOfLastUpdate int64
	Status                TaskStatus
	OpenedByUserID        uint
	Kind                  TaskKind
}

type TaskInfo struct {
	Name string

	TaskMetadata
}

type Task struct {
	PrimaryKeyTask

	TaskInfo
}
