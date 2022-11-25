package resource

import (
	"fmt"
	"fyne.io/fyne/v2"
	"github.com/konglingyinxia/go-jar-encryption/logger"
	"io/fs"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

const IconPath = "resource/img/favicon.ico"
const fontPath = "resource/font"

var fontFiles []string

func init() {
	err := filepath.Walk(fontPath, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			fontFiles = append(fontFiles, path)
		}
		return err
	})
	if err != nil {
		logger.Log().Error("初始化字体库失败", err)
	}
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(fontFiles))
	path := fontFiles[i]
	err = os.Setenv("FYNE_FONT", path)
	if err != nil {
		logger.Log().Error("初始化字体库环境变量设置失败", err)
	}
}

func LoadResource(path string) fyne.Resource {
	r, err := fyne.LoadResourceFromPath(path)
	if err != nil {
		fmt.Printf("%s,图片资源加载失败.....\n", path)
	}
	return r
}
