package cron

import (
	"time"
)

// DelaySchedule represents once task.
type DelaySchedule struct {
	runTime time.Time
}

// Delay func returns a crontab Schedule that activates delay duration.
// Delays of less than a second are not supported (will panic).
// Any fields less than a Second are truncated.
func Delay(delay time.Duration) DelaySchedule {
	if delay < time.Second {
		panic("cron/constantdelay: delays of less than a second are not supported: " +
			delay.String())
	}
	now := time.Now()
	return DelaySchedule{
		runTime: now.Add(delay),
	}
}

// RunAt func returns a crontab Schedule that activates at the target time.
func RunAt(runtime time.Time) DelaySchedule {
	now := time.Now()
	if now.After(runtime) {
		panic("runtime need after now ")
	}
	return DelaySchedule{runTime: runtime}
}


// Next returns the next time this should be run.
// This rounds so that the next activation time will be on the second.
func (schedule DelaySchedule) Next(t time.Time) time.Time {
	if t.Before(schedule.runTime) {
		return time.Time{}
	}
	return schedule.runTime
}
