package main

import (
	"fmt"
	"time"

	s "github.com/washed/kitchen-sink-go/schedules"
)

func main() {
	start, _ := time.Parse(time.RFC3339, "2023-02-16T01:00:00+01:00")
	end, _ := time.Parse(time.RFC3339, "2023-02-16T01:02:00+01:00")

	scheduleNow := s.AbsoluteSchedule{
		Start: start,
		End:   end,
	}

	scheduleFuture := s.AbsoluteSchedule{
		Start: start.Add(2 * time.Minute),
		End:   end.Add(2 * time.Minute),
	}

	schedulePast := s.AbsoluteSchedule{
		Start: start.Add(-2 * time.Minute),
		End:   end.Add(-2 * time.Minute),
	}

	schedules := []s.Schedule{
		schedulePast,
		scheduleNow,
		scheduleFuture,
	}

	for _, schedule := range schedules {
		fmt.Printf("Schedule is active: %t\n", schedule.IsActive())
		fmt.Printf("time until start: %s\n", schedule.GetTimeUntilStart())
		fmt.Printf("time until end: %s\n\n", schedule.GetTimeUntilEnd())
	}

	activeSchedules := s.GetActiveSchedules(schedules)

	fmt.Printf("active schedules: %s\n", activeSchedules)
}
