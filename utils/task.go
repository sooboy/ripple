package utils

import (
	"time"
)

//Task 简单的任务
type Task struct {
	ticker   *time.Ticker
	callback func()
}

//init 初始化任务
func (t *Task) init() {
	for _ = range t.ticker.C {
		t.callback()
	}
}
func NewTask(tickerTime int64, callback func()) *Task {
	t := &Task{
		ticker:   time.NewTicker(time.Duration(tickerTime) * time.Second),
		callback: callback,
	}
	go t.init()
	return t
}
