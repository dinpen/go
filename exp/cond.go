package exp

// 三元运算函数
func If[T any](b bool, v1, v2 T) T {
	if b {
		return v1
	}
	return v2
}

// 返回一个非零值
func ValOrDefault[T comparable](v T, def T) T {
	var zero T
	if v == zero {
		return def
	}
	return v
}
