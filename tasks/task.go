package tasks

import (
	"runtime"
	"time"
)

type Task interface {
	Do()
}

type taskQueue chan Task

type worker struct {
	tq taskQueue
}

func newWorker() worker {
	return worker{tq: make(taskQueue)}
}

func (w worker) work(wq chan taskQueue) {
	go func() {
		for {
			wq <- w.tq
			job := <-w.tq
			job.Do()
		}
	}()
}

type taskPool struct {
	routines int
	tq       taskQueue
	wq       chan taskQueue
}

func NewTaskPool(routines int) *taskPool {
	return &taskPool{
		routines: routines,
		tq:       make(taskQueue),
		wq:       make(chan taskQueue, routines),
	}
}
func (pool *taskPool) Open() {
	for i := 0; i < pool.routines; i++ {
		worker := newWorker()
		worker.work(pool.wq)
	}

	go func() {
		for {
			task := <-pool.tq
			// Acquire a worker.
			worker := <-pool.wq
			// Assign task to a worker's task queue.
			worker <- task
		}
	}()
}

func (pool *taskPool) Accept(t Task) {
	pool.tq <- t
}

func (pool *taskPool) Close() {
	pool = nil
}

func DoAfter(d time.Duration, f func()) {
	timer := time.AfterFunc(d, f)
	defer timer.Stop()

	<-timer.C
}

func DoWithTimeout(d time.Duration, f func()) bool {
	timeout := time.NewTimer(d)
	defer timeout.Stop()

	done := make(chan struct{}, 1)
	go func() {
		f()
		done <- struct{}{}
	}()

	select {
	case <-timeout.C:
		return false
	case <-done:
		close(done)
		return true
	}
}

type intervalOption struct {
	interval       time.Duration
	maxConcurrency int
}

func newIntervalOption(interval time.Duration, maxConcurrency int) *intervalOption {
	return &intervalOption{
		interval:       interval,
		maxConcurrency: maxConcurrency,
	}
}

func (o *intervalOption) Interval() time.Duration {
	return o.interval
}

func (o *intervalOption) MaxConcurrency() int {
	if o.maxConcurrency <= 0 {
		return runtime.GOMAXPROCS(0)
	}
	return o.maxConcurrency
}

func DoInterval(opt intervalOption, f func()) {
	timer := time.NewTicker(opt.Interval())
	defer timer.Stop()
	next := make(chan struct{}, opt.MaxConcurrency())

	for {
		<-timer.C
		next <- struct{}{}
		go func() {
			f()
			<-next
		}()
	}
}
