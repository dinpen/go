package times

import "time"

// 秒转换为时间
func TimeFromSeconds(v int64) time.Time {
	return time.Unix(v/1000, 0)
}

// 毫秒转换为时间
func TimeFromMicroseconds(v int64) time.Time {
	return time.Unix(0, v*int64(time.Microsecond))
}

// 微秒转换为时间
func TimeFromMilliseconds(v int64) time.Time {
	return time.Unix(0, v*int64(time.Millisecond))
}

// 时间转换为微秒
func MillisecondsFromTime(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

// 时间转换为毫秒
func MicrosecondsFromTime(t time.Time) int64 {
	return t.UnixNano() / int64(time.Microsecond)
}
