package maps

// 合并 map
func Merge[K comparable, V any](maps ...map[K]V) map[K]V {
	result := make(map[K]V)
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}

// 获取 map 的 key 列表
func Keys[K comparable, V any](m map[K]V) []K {
	result := make([]K, len(m))
	for k := range m {
		result = append(result, k)
	}
	return result
}

// 获取 map 的 value 列表
func Values[K comparable, V any](m map[K]V) []V {
	result := make([]V, len(m))
	for _, v := range m {
		result = append(result, v)
	}
	return result
}

// 判断 map 是否包含指定的 key
func Contains[K comparable, V any](m map[K]V, k K) bool {
	_, ok := m[k]
	return ok
}

// 判断 map 是否包含指定的 value
func ContainsValue[K comparable, V comparable](m map[K]V, v V) bool {
	for _, i := range m {
		if v == i {
			return true
		}
	}
	return false
}

// 过滤 map
func Filter[K comparable, V any](m map[K]V, f func(k K, v V) bool) map[K]V {
	result := make(map[K]V)
	for k, v := range m {
		if f(k, v) {
			result[k] = v
		}
	}
	return result
}

// 生成新的 map
func Remap[K comparable, V any, U any](m map[K]V, f func(k K, v V) U) map[K]U {
	result := make(map[K]U)
	for k, v := range m {
		result[k] = f(k, v)
	}
	return result
}
