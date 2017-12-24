package model

import (
	"time"
)

type Task struct {
	ID          string
	ScheduleID  string
	UserID      string
	StartTime   time.Time
	Duration    int64
	SignTimes   int
	CreatedTime time.Time
	UpdatedTime time.Time
}
