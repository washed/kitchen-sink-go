package main

import (
	"time"

	s "github.com/washed/kitchen-sink-go/schedules"
)

func main() {
	start, _ := time.Parse(time.RFC3339, "2023-02-16T01:00:00+01:00")
	end, _ := time.Parse(time.RFC3339, "2023-02-16T01:02:00+01:00")

	schedules := []s.AbsoluteSchedule{
		{
			Start: start,
			End:   end,
		}}

	s.GetTopActiveSchedule(schedules)
}
