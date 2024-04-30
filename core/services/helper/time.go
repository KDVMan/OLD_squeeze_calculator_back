package core_services_helper

import (
	"time"
)

func MillisecondsToTime(milliseconds int64) string {
	t := time.UnixMilli(milliseconds)

	return t.Format("02.01.2006 15:04:05")
}
