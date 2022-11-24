package resource

import (
	"fmt"
	"fyne.io/fyne/v2"
)

const IconPath = "resource/img/favicon.ico"

func LoadResource(path string) fyne.Resource {
	r, err := fyne.LoadResourceFromPath(path)
	if err != nil {
		fmt.Printf("%s,图片资源加载失败.....\n", path)
	}
	return r
}
