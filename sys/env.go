package sys

import "os"

// 获取环境变量的值，如果不存在则返回默认值
func GetEnvOrDefault(key, defaultVal string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultVal
}
