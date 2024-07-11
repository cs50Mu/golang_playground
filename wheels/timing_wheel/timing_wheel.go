package timing_wheel

import (
	"fmt"
	"time"
)

type rawTime struct {
	Day  int
	Hour int
	Min  int
	Sec  int
}

type Task struct {
	Name     string
	execFunc func()
	// time to exec
	execTime rawTime
}

type TimingWheel struct {
	Day  *timingWheel
	Hour *timingWheel
	Min  *timingWheel
	Sec  *timingWheel
}

type wheelType int

const (
	wheelTypeDay  wheelType = 1
	wheelTypeHour           = 2
	wheelTypeMin            = 3
	wheelTypeSec            = 4
)

type timingWheel struct {
	// the elapsed time for every tick
	tickDur time.Duration
	// the underlying array
	arr []*DoublyLinkedList[*Task, int]
	// length of the array
	len int64
	// the idx where the pointer points currently
	currIdx   int64
	wheelType wheelType
}

func reInsertTask(sourceTw, targetTw *timingWheel) {
	getExecTime := func(task *Task) int {
		var res int
		switch targetTw.wheelType {
		case wheelTypeHour:
			res = task.execTime.Hour
		case wheelTypeMin:
			res = task.execTime.Min
		case wheelTypeSec:
			res = task.execTime.Sec
		}
		return res
	}
	for _, task := range sourceTw.arr[sourceTw.currIdx].Keys() {
		targetIdx := (targetTw.currIdx + int64(getExecTime(task))) % targetTw.len
		fmt.Printf("insert task: %+v to targetIdx: %v\n", task, targetIdx)
		fmt.Printf("%v --> %v\n", sourceTw.wheelType, targetTw.wheelType)
		targetTw.arr[targetIdx].PushBack(task, 0)
	}
	// empty it
	sourceTw.arr[sourceTw.currIdx] = NewDoublyLinkedList[*Task, int]()
}

// func (tw *timingWheel) tick() {
// 	tw.currIdx++
// 	if tw.arr[tw.currIdx].Len > 0 {
// 		switch tw.wheelType {
// 		case wheelTypeDay:
// 			// reInsertTask(tw.)
// 		}
// 	}
// }

func newTW(tickDur time.Duration, len int64, wheelType wheelType) *timingWheel {
	arr := make([]*DoublyLinkedList[*Task, int], len)
	for i := int64(0); i < len; i++ {
		arr[i] = NewDoublyLinkedList[*Task, int]()
	}
	return &timingWheel{
		tickDur:   tickDur,
		arr:       arr,
		len:       len,
		wheelType: wheelType,
	}
}

func NewTimingWheel() *TimingWheel {
	dayWheel := newTW(24*time.Hour, 100, wheelTypeDay)
	hourWheel := newTW(1*time.Hour, 24, wheelTypeHour)
	minWheel := newTW(1*time.Minute, 60, wheelTypeMin)
	secWheel := newTW(1*time.Second, 60, wheelTypeSec)
	return &TimingWheel{
		Day:  dayWheel,
		Hour: hourWheel,
		Min:  minWheel,
		Sec:  secWheel,
	}
}

// func (tw *TimingWheel) tickSec() {
// 	tw.Sec.currIdx++

// }

// func (tw *TimingWheel) tickMin() {

// }

// func (tw *TimingWheel) tickHour() {

// }

func (tw *TimingWheel) Start() {
	for {
		if tw.Sec.arr[tw.Sec.currIdx].Len > 0 {
			tq := tw.Sec.arr[tw.Sec.currIdx]
			curr := tq.Head.Next
			for curr != tq.Tail {
				// remove it && execute it
				tq.Remove(curr)
				curr.Key.execFunc()
				curr = curr.Next
			}
		}
		tw.Sec.currIdx++
		fmt.Printf("sec: %v, min: %v\n", tw.Sec.currIdx, tw.Min.currIdx)
		if tw.Sec.currIdx == tw.Sec.len {
			fmt.Println("reset sec to 0")
			tw.Sec.currIdx = 0
			tw.Min.currIdx++
			if tw.Min.arr[tw.Min.currIdx].Len > 0 {
				fmt.Println("before reinsert task to Sec")
				reInsertTask(tw.Min, tw.Sec)
				fmt.Println("reinsert task to Sec done")
			}
			if tw.Min.currIdx == tw.Min.len {
				tw.Min.currIdx = 0
				tw.Hour.currIdx++
				if tw.Hour.arr[tw.Hour.currIdx].Len > 0 {
					reInsertTask(tw.Hour, tw.Min)
				}
				if tw.Hour.currIdx == tw.Hour.len {
					tw.Hour.currIdx = 0
					tw.Day.currIdx++
					if tw.Day.arr[tw.Day.currIdx].Len > 0 {
						reInsertTask(tw.Day, tw.Hour)
					}
					if tw.Day.currIdx == tw.Day.len {
						tw.Day.currIdx = 0
					}
				}
			}
		}
		time.Sleep(tw.Sec.tickDur)
	}
}

func calcTicks(delay time.Duration, tickDur time.Duration) (ticks int, rem time.Duration) {
	return int(delay / tickDur), delay % tickDur
}

func setTaskTime(task *Task, delay time.Duration) {
	var rem time.Duration
	task.execTime.Day, rem = calcTicks(delay, 24*time.Hour)
	task.execTime.Hour, rem = calcTicks(rem, time.Hour)
	task.execTime.Min, rem = calcTicks(rem, time.Minute)
	task.execTime.Sec, rem = calcTicks(rem, time.Second)
}

func (tw *TimingWheel) now() (time.Duration, *rawTime) {
	nowTime := &rawTime{
		Day:  int(tw.Day.currIdx),
		Hour: int(tw.Hour.currIdx),
		Min:  int(tw.Min.currIdx),
		Sec:  int(tw.Sec.currIdx),
	}
	return time.Duration(tw.Day.currIdx)*time.Hour*24 +
		time.Duration(tw.Hour.currIdx)*time.Hour +
		time.Duration(tw.Min.currIdx)*time.Minute +
		time.Duration(tw.Sec.currIdx)*time.Second, nowTime
}

func (tw *TimingWheel) SubmitTask(task *Task, delay time.Duration) {
	nowDur, nowTime := tw.now()
	futureTime := nowDur + delay
	setTaskTime(task, futureTime)

	fmt.Printf("now: %+v, taskTime: %+v\n", nowTime, task.execTime)

	if task.execTime.Day > nowTime.Day {
		tw.Day.arr[task.execTime.Day].PushBack(task, 0)
	} else if task.execTime.Hour > nowTime.Hour {
		tw.Hour.arr[task.execTime.Hour].PushBack(task, 0)
	} else if task.execTime.Min > nowTime.Min {
		tw.Min.arr[task.execTime.Min].PushBack(task, 0)
	} else {
		tw.Sec.arr[task.execTime.Sec].PushBack(task, 0)
	}
}
