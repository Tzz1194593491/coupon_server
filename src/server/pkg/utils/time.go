package utils

import "time"

const (
	TEMPLATE = "2006-01-02 15:04:05"
)

func StringToTime(needConvertTime string) (time.Time, error) {
	return time.Parse(TEMPLATE, needConvertTime)
}

func TimeToString(needConvertTime time.Time) string {
	return needConvertTime.Format(TEMPLATE)
}
