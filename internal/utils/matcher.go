package utils

import (
	"time"
)

func MatchSchedule(schedule string, t time.Time) bool {
	switch schedule {
	case "* * * * *":
		return true
	case t.Format("15:04"):
		return true
	default:
		return false
	}
}
