package schedules

import (
	"errors"
	"time"
)

type Schedule interface {
	IsActive() bool
	GetTimeUntilStart() time.Duration
	GetTimeUntilEnd() time.Duration
}

type WeekdaySchedule struct {
	Weekdays []time.Weekday `json:"weekdays" yaml:"weekdays"`
	Start    time.Duration  `json:"start"    yaml:"start"`
	End      time.Duration  `json:"end"      yaml:"end"`
}

type AbsoluteSchedule struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

var location *time.Location = time.UTC

func SetLocation(locationName string) error {
	loc, err := time.LoadLocation(locationName)
	if err != nil {
		return err
	}

	location = loc
	return nil
}

func GetNow() time.Time {
	return time.Now().In(location)
}

func (s WeekdaySchedule) IsActiveToday() bool {
	weekdayToday := GetNow().Weekday()

	for _, v := range s.Weekdays {
		if v == weekdayToday {
			return true
		}
	}

	return false
}

func (s WeekdaySchedule) NextExecutionDay() time.Weekday {
	weekdayToday := GetNow().Weekday()

	for _, v := range s.Weekdays {
		if v == weekdayToday {
			return true
		}
	}

	return false
}

func (s WeekdaySchedule) ToAbsoluteSchedule() (*AbsoluteSchedule, error) {
	if !s.IsActiveToday() {
		return nil, errors.New("schedule is not active today")
	}

	today0 := GetNow().Truncate(24 * time.Hour)

	absoluteSchedule := AbsoluteSchedule{
		Start: today0.Add(s.Start),
		End:   today0.Add(s.End),
	}

	return &absoluteSchedule, nil
}

func (s WeekdaySchedule) IsActive() bool {
	absoluteSchedule, err := s.ToAbsoluteSchedule()
	if err != nil {
		return false
	}
	return absoluteSchedule.IsActive()
}

func (s WeekdaySchedule) GetTimeUntilStart() time.Duration {
	return s.Start.Sub(GetNow())
}

func (s WeekdaySchedule) GetTimeUntilEnd() time.Duration {
	return s.End.Sub(GetNow())
}

func (s AbsoluteSchedule) IsActive() bool {
	nowLocal := GetNow()
	return s.Start.Before(nowLocal) && s.End.After(nowLocal)
}

func (s AbsoluteSchedule) GetTimeUntilStart() time.Duration {
	return s.Start.Sub(GetNow())
}

func (s AbsoluteSchedule) GetTimeUntilEnd() time.Duration {
	return s.End.Sub(GetNow())
}

func GetActiveSchedules(schedules []Schedule) []Schedule {
	var _ Schedule = AbsoluteSchedule{}

	var activeSchedules []Schedule
	for _, s := range schedules {
		if s.IsActive() {
			activeSchedules = append(activeSchedules, s)
		}
	}

	return activeSchedules
}

func GetTopActiveSchedule(schedules []Schedule) Schedule {
	return GetActiveSchedules(schedules)[0]
}
