package strings

import (
	"unicode/utf8"
)

// 字符串的字节长度
func BytesCount(s string) int {
	return len(s)
}

// 字符串的字符长度
func CharsCount(s string) int {
	return utf8.RuneCountInString(s)
}

// 缩短字符串
func ShortenString(v string, n int) string {
	if len(v) <= n {
		return v
	}
	return v[:n]
}
