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

// 返回 panic 信息
func (p *recoveredPanic) PanicInfo() string {
	return fmt.Sprintf("panic: %v\nstacktrace: \n%s\n", p.val, string(p.stack))
}

// 返回 panic 的值
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

// WaitGroup 是 sync.WaitGroup 的扩展，可以捕获 goroutine 中的 panic
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

// Run 会并发执行 f 中的函数
func Run(f ...gofunc) {
	for _, v := range f {
		go v()
	}
}

// Wait 会并发执行 f 中的函数，并等待所有函数执行完毕
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

// WaitAndRecover 会并发执行 f 中的函数，并等待所有函数执行完毕
// 如果有 goroutine panic，会返回 panic 信息
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

// ForIndex 会并发执行 f 中的函数 n 次，并等待执行完毕
func ForIndex(n int, f func(i int)) {
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(i int) {
			defer wg.Done()
			f(i)
		}(i)
	}
	wg.Wait()
}

// ForEach 会并发执行 f 中的函数并等待所有函数执行完毕
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

// ForEachWithLimit 会限制并发次数执行 f 中的函数并等待所有函数执行完毕
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
