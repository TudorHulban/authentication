package helpers

import (
	"database/sql"
	"time"
)

type Timestamp struct {
	CreatedAt int64         `json:",omitempty"`
	UpdatedAt int64         `json:",omitempty"`
	DeletedAt sql.NullInt64 `json:",omitempty"`
}

func (t *Timestamp) WithCreateNow() Timestamp {
	t.CreatedAt = time.Now().UnixNano()

	return *t
}

func (t *Timestamp) WithUpdateNow() Timestamp {
	t.UpdatedAt = time.Now().UnixNano()

	return *t
}

func (t *Timestamp) WithDeleteNow() Timestamp {
	t.DeletedAt = sql.NullInt64{
		Valid: true,
		Int64: time.Now().UnixNano(),
	}

	return *t
}
