package task

type TaskStatus uint8

const (
	StatusNew        = TaskStatus(0)
	StatusInProgress = TaskStatus(1)
	StatusClosed     = TaskStatus(2)
)
