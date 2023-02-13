package file

import (
	"os"
	"runtime"
)

// 判断文件或文件夹是否存在
func IsPathExist(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// 获取用户主目录
func HomePath() string {
	res := os.Getenv("HOME")
	if runtime.GOOS == "windows" {
		if len(res) > 0 {
			return res
		}

		if res = os.Getenv("USERPROFILE"); len(res) > 0 {
			return res
		}

		if homeDrive, homePath := os.Getenv("HOMEDRIVE"), os.Getenv("HOMEPATH"); len(homeDrive) > 0 && len(homePath) > 0 {
			res = homeDrive + homePath
			return res
		}
	}
	return res
}
