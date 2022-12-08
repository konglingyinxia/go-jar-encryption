package projectpath

import (
	"github.com/konglingyinxia/go-jar-encryption/logger"
	"os"
	"path"
	"runtime"
)

func RootPath() string {
	s, _ := os.Getwd()
	return s
}
func GetWorkPath() string {
	_, filename, _, _ := runtime.Caller(0)
	logger.Log().Info("RootPath:", filename)
	root := path.Dir(path.Dir(filename))
	return root
}
