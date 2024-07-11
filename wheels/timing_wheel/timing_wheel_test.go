package timing_wheel

import (
	"fmt"
	"testing"
	"time"
)

func TestTimingWheel(t *testing.T) {
	tw := NewTimingWheel()

	tw.SubmitTask(newTask("task1"), 5*time.Second)
	tw.SubmitTask(newTask("task2"), 65*time.Second)
	tw.SubmitTask(newTask("task2"), 77*time.Second)

	tw.Start()

	// time.Sleep(60 * time.Second)
}

func newTask(name string) *Task {
	return &Task{
		Name:     name,
		execFunc: func() { fmt.Printf("exec task: %v at %+v\n", name, time.Now()) },
	}
}

func TestCalcTicks(t *testing.T) {
	task := newTask("dummy")
	now := 11*24*time.Hour + 10*time.Hour + 24*time.Minute + 30*time.Second
	delay := 50*time.Minute + 45*time.Second
	setTaskTime(task, now+delay)

	fmt.Printf("task: %+v\n", task)
}
