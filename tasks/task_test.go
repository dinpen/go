package tasks

import (
	"fmt"
	"log"
	"testing"
	"time"
)

type Alarm struct {
	Message string
}

func (a Alarm) Do() {
	fmt.Println("ding ding ding ~")
	fmt.Println(a.Message)
}

func TestTask(t *testing.T) {
	taskPool := NewTaskPool(5)
	taskPool.Open()

	go func() {
		for i := 0; i < 20; i++ {
			taskPool.Accept(Alarm{
				Message: fmt.Sprintf("[%d] it's time to exercise.", i),
			})
			//time.Sleep(time.Second)
		}
	}()

	// for {
	// 	fmt.Printf("runtime.NumGoroutine: %d\n", runtime.NumGoroutine())
	// }
}

func TestDoInterval(t *testing.T) {
	DoInterval(*newIntervalOption(time.Second, 5), func() {
		log.Println("start1")
		time.Sleep(time.Second * 10)
		log.Println("end1")
	})

	DoInterval(*newIntervalOption(4*time.Second, 5), func() {
		fmt.Println("start2")
		time.Sleep(time.Second)
		fmt.Println("end2")
	})
}

func TestDoWithTimeout(t *testing.T) {
	done1 := DoWithTimeout(time.Second*2, func() {
		fmt.Println("start1")
		fmt.Println("end1")
	})
	fmt.Printf("done1: %v\n", done1)

	done2 := DoWithTimeout(time.Second*2, func() {
		fmt.Println("start2")
		time.Sleep(time.Second * 5)
		fmt.Println("end2")
	})
	fmt.Printf("done2: %v\n", done2)
}

func TestDoAfter(t *testing.T) {
	log.Println("test start")
	DoAfter(time.Second*2, func() {
		log.Println("do start")
	})
}
