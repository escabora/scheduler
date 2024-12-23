package models

import "time"

type Task struct {
	Schedule string
	Job      func()
}

type PredictedTask struct {
	Name      string
	NextRun   time.Time
	Job       func()
	LastRun   time.Time
	Frequency time.Duration
}
