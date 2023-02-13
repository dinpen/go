package tasks

// 运行函数返回 false 时，重试
func RetryOnFalse(times int, f func() bool) {
	for i := 0; i < times; i++ {
		if f() {
			return
		}
	}
}

// 运行函数 error 时，重试
func RetryOnError(times int, f func() error) {
	for i := 0; i < times; i++ {
		if err := f(); err == nil {
			return
		}
	}
}
