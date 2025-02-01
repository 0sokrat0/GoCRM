package entity

import (
	"errors"
	"time"
)

type TimeInterval struct {
	Start time.Time `json:"start_time"`
	End   time.Time `json:"end_time"`
}

func (ti TimeInterval) Validate() error {
	if !ti.End.After(ti.Start) {
		return errors.New("end time must be after start time")
	}
	return nil
}

type WorkingSchedule struct {
	DayOfWeek int            `json:"day_of_week"`      /// 0 = Воскресенье, 6 = Суббота
	WorkTime  TimeInterval   `json:"work_time"`        // Рабочее время (например, 09:00-17:00)
	Breaks    []TimeInterval `json:"breaks,omitempty"` // Опционально: перерывы в рабочее время
}
