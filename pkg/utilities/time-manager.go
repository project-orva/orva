package utilities

import "time"

type TimeStamp struct {
	StartTime int64
	EndTime   int64
}

func (t *TimeStamp) Start() {
	t.StartTime = time.Now().Unix()
}

func (t *TimeStamp) Stop() {
	t.EndTime = time.Now().Unix()
}

func (t *TimeStamp) Distance() int64 {
	return t.EndTime - t.StartTime
}
