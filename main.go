package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/konglingyinxia/go-jar-encryption/layout"
	"github.com/konglingyinxia/go-jar-encryption/logger"
	"github.com/konglingyinxia/go-jar-encryption/resource"
)

var win fyne.Window

func start() {
	a := app.NewWithID("com.mzydz.jarencryption")
	win = a.NewWindow("jar包加密")
	win.Resize(fyne.NewSize(800, 500))
	win.SetIcon(resource.LoadResource(resource.IconPath))

	layout.BaseFrom(win)
	win.SetFixedSize(true)
	logger.Log().Info("启动完成.............")
	win.SetMaster()
	// 窗口居中
	win.CenterOnScreen()
	win.ShowAndRun()
	logger.Log().Info("服务退出.............")
}

func main() {
	start()
	//xjarGoPath := "/home/kongling/桌面/test/ad_publish_bank_test/out/xjar.go"
	//dir := filepath.Dir(xjarGoPath)
	//os2.Setenv("GOARCH", "amd64")
	//os2.Setenv("GOOS", "linux")
	/////home/kongling/work/work/project/go/my/go-java-jar-encryption/env/go/go_linux/bin/go
	//cmdXjarGo := exec.Command("go", "build", xjarGoPath)
	//cmdXjarGo.Dir = dir
	//var out bytes.Buffer
	//var stderr bytes.Buffer
	//cmdXjarGo.Stdout = &out
	//cmdXjarGo.Stderr = &stderr
	//err := cmdXjarGo.Run()
	//if err != nil {
	//	log.Println(err.Error(), stderr.String())
	//} else {
	//	log.Println(out.String())
	//}
}
