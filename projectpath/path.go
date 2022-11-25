package projectpath

import (
	"path"
	"runtime"
)

func RootPath() string {
	_, filename, _, _ := runtime.Caller(0)
	root := path.Dir(path.Dir(filename))
	return root
}
