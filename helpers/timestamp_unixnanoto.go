package helpers

import (
	"time"
)

func UnixNanoTo(timestamp int64) string {
	return time.Unix(0, timestamp).
		Format("2006, Jan 02 15:04:05")
}
