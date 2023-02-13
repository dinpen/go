package slices

import (
	"math/rand"
	"time"
)

// 合并 slice
func Merge[T any](items ...[]T) []T {
	result := make([]T, 0)
	for _, m := range items {
		result = append(result, m...)
	}
	return result
}

// 判断 slice 是否包含指定的元素
func Contains[T comparable](items []T, v T) bool {
	for _, i := range items {
		if v == i {
			return true
		}
	}
	return false
}

// 去重
func Distinct[T comparable](items []T) []T {
	keys := map[T]bool{}
	for _, k := range items {
		keys[k] = true
	}
	result := make([]T, len(keys))
	for k := range keys {
		result = append(result, k)
	}
	return result
}

// 随机排序
func Shuffle[T any](items []T) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := len(items) - 1; i > 0; i-- {
		j := random.Intn(i + 1)
		items[i], items[j] = items[j], items[i]
	}
}

// 反转
func Reverse[T any](items []T) {
	l := len(items)
	for i := l/2 - 1; i >= 0; i-- {
		items[i], items[l-i-1] = items[l-i-1], items[i]
	}
}

// 生成新的 slice
func Reslice[T, U any](items []T, f func(t T) U) []U {
	res := make([]U, len(items))
	for i := range items {
		res[i] = f(items[i])
	}
	return res
}

// 过滤
func Filter[T any](items []T, f func(t T) bool) []T {
	res := make([]T, 0)
	for _, i := range items {
		if f(i) {
			res = append(res, i)
		}
	}
	return res
}
