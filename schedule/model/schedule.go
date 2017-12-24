package model

import (
	"time"
)

type Schedule struct {
	ID          string
	Name        string
	Description string
	Period      int64
	CreatedTime time.Time
	UpdatedTime time.Time
}

type ScheduleFact struct {
	ID         string
	ScheduleID string
	Name       string
	Value      string
}
