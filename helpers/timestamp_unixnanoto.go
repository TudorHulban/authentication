package helpers

import (
	"time"
)

func UnixNanoTo(timestamp int64) string {
	return time.Unix(0, timestamp).
		Format("2006, January 02 HH15:04:05")
}
