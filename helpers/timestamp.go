package helpers

import (
	"database/sql"
	"fmt"
	"strconv"
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

func UnixNanoTo(timestamp int64) string {
	then := time.Unix(0, timestamp)

	y, m, d := then.Date()

	return fmt.Sprintf(
		"%s, %s %d HH%d:%d:%d",
		strconv.Itoa(y),
		m.String(),
		d,
		then.Hour(),
		then.Minute(),
		then.Second(),
	)
}
