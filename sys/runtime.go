package sys

import (
	"runtime"
	"strings"
)

// 获取当前函数的完整名称
func CurrentFuncFullName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	return runtime.FuncForPC(pc[0]).Name()
}

// 获取当前函数的名称
func CurrentFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)

	fullName := runtime.FuncForPC(pc[0]).Name()
	strs := strings.Split(fullName, ".")
	return strs[len(strs)-1]
}
