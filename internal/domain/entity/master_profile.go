package entity

import (
	"time"

	"github.com/google/uuid"
)

type MasterProfile struct {
	MasterID uuid.UUID         `json:"master_id"`
	Schedule []WorkingSchedule `json:"schedule"`
}

func (m *MasterProfile) IsAvailable(bookingTime time.Time, duration time.Duration) bool {
	bookingEnd := bookingTime.Add(duration)

	bookingDay := int(bookingTime.Weekday())

	for _, ws := range m.Schedule {
		if ws.DayOfWeek != bookingDay {
			continue
		}
		if bookingTime.After(ws.WorkTime.Start) && bookingEnd.Before(ws.WorkTime.End) {
			available := true
			for _, br := range ws.Breaks {
				if bookingTime.Before(br.End) && bookingEnd.After(br.Start) {
					available = false
					break
				}
			}
			if available {
				return true
			}
		}
	}
	return false
}
