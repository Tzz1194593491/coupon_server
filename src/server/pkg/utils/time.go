package utils

import "time"

const (
	TEMPLATE = "2006-01-02 15:04:05"
)

func UnixSecondToTime(second int64) time.Time {
	return time.Unix(second, 0)
}
