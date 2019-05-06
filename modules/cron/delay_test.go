package cron

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRunAt(t *testing.T) {
	now := time.Now().Add(time.Second)
	s := RunAt(now)
	p := DelaySchedule{runTime:now}
	fmt.Println(s.runTime, p.runTime)
	ss := s.runTime.Format("2006-01-02 15:04:05")
	sp := p.runTime.Format("2006-01-02 15:04:05")
	assert.Equal(t, ss, sp, "need return DelaySchedule")
}

func TestDelaySchedule_Next(t *testing.T) {
	now := time.Now()
	yes := time.Now().Add(time.Second * 10)
	d := DelaySchedule{runTime:now}
	y := DelaySchedule{runTime:yes}
	dn := d.runTime
	yn := d.runTime
	assert.Equal(t, d.Next(now), dn, "error next time")
	assert.NotEqual(t, y.Next(now), yn, "error next time")
}