package task

type TaskKind uint8

const (
	KindUndefined = TaskKind(0)
	KindTicket    = TaskStatus(1)
)
