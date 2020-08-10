package timer

import (
	"os/exec"
	"sync/atomic"
	"time"
)

type TimerConf struct {
	Interval   time.Duration
	Command    string
	Args       []string
	TotalTimes int
}

func (t *TimerConf) Exec() {
	cmd := exec.Command(t.Command, t.Args...)

	execFunc := func(b *int32) {
		cmd.Run()
		atomic.StoreInt32(b, 1)
	}

	for times := 0; times < t.TotalTimes; times++ {
		tr := time.NewTimer(t.Interval)
		flagInt := int32(0)

		go execFunc(&flagInt)
		select {
		case <-tr.C:
			if atomic.LoadInt32(&flagInt) != 1 {
				panic("command is not done in timeslot")
			}
		}
	}
}
