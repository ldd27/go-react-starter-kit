package tool

import (
	"os"
)

//检查目录是否存在
func CheckFileIsExist(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}
