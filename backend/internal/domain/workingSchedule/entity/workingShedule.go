package entity

import (
	"time"

	"github.com/google/uuid"
)

type WorkingSchedule struct {
	ID        uuid.UUID   `json:"id"`
	MasterID  uuid.UUID   `json:"master_id"`
	Day       int         `json:"day"`
	StartTime time.Time   `json:"start_time"`
	EndTime   time.Time   `json:"end_time"`
	IsDayOff  bool        `json:"is_day_off"`
	Breaks    []BreakTime `json:"breaks"`
}

type BreakTime struct {
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}
