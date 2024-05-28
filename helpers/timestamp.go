package helpers

import (
	"fmt"
	"strconv"
	"time"
)

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
