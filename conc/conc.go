package conc

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"sync"
	"sync/atomic"
)

type gofunc func()

type recoveredPanic struct {
	val     any
	callers []uintptr
	stack   []byte
}

func (p *recoveredPanic) PanicInfo() string {
	return fmt.Sprintf("panic: %v\nstacktrace: \n%s\n", p.val, string(p.stack))
}

func (p *recoveredPanic) PanicVal() any {
	return p.val
}

func newRecoveredPanic(v any) *recoveredPanic {
	var callers []uintptr
	n := runtime.Callers(2, callers)
	return &recoveredPanic{
		val:     v,
		callers: callers[:n],
		stack:   debug.Stack(),
	}
}

type WaitGroup struct {
	sync.WaitGroup
	recovered atomic.Pointer[recoveredPanic]
}

func (wg *WaitGroup) try(f gofunc) {
	defer func() {
		if r := recover(); r != nil {
			wg.recovered.CompareAndSwap(nil, newRecoveredPanic(r))
		}
	}()

	f()
}

func Run(f ...gofunc) {
	for _, v := range f {
		go v()
	}
}

func Wait(f ...gofunc) {
	var wg sync.WaitGroup
	wg.Add(len(f))
	for _, v := range f {
		go func(v gofunc) {
			defer wg.Done()
			v()
		}(v)
	}
	wg.Wait()
}

func WaitAndRecover(f ...gofunc) *recoveredPanic {
	var wg WaitGroup
	wg.Add(len(f))
	for _, v := range f {
		go func(v gofunc) {
			defer wg.Done()
			wg.try(v)
		}(v)
	}
	wg.Wait()
	return wg.recovered.Load()
}

func ForIndex(num int, f func(i int)) {
	var wg sync.WaitGroup
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func(i int) {
			defer wg.Done()
			f(i)
		}(i)
	}
	wg.Wait()
}

func ForEach[T any](items []T, f func(item T)) {
	var wg sync.WaitGroup
	wg.Add(len(items))
	for _, item := range items {
		go func(item T) {
			defer wg.Done()
			f(item)
		}(item)
	}
	wg.Wait()
}

func ForEachWithLimit[T any](items []T, f func(item T), limit int) {
	if limit == 0 {
		limit = runtime.GOMAXPROCS(0)
	}
	l := len(items)
	if limit > l {
		limit = l
	}

	limiter := make(chan struct{}, limit)
	for i := 0; i < limit; i++ {
		limiter <- struct{}{}
	}

	tasks := make(chan T, l)
	for _, item := range items {
		tasks <- item
	}
	close(tasks)

	var wg sync.WaitGroup
	wg.Add(l)
	for t := range tasks {
		<-limiter
		go func(t T) {
			defer wg.Done()
			f(t)
			limiter <- struct{}{}
		}(t)
	}
	wg.Wait()
	close(limiter)
}
