package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/konglingyinxia/go-jar-encryption/layout"
	"github.com/konglingyinxia/go-jar-encryption/logger"
	"github.com/konglingyinxia/go-jar-encryption/resource"
	"os"
	"runtime"
)

var win fyne.Window

func init() {
	//初始化系统环境变量
	sysType := runtime.GOOS
	if sysType == "linux" {
		//java环境 jdk1.8_352
		os.Setenv("JAVA_HOME", "env/jvm/jre_linux")
		os.Setenv("PATH", "$JAVA_HOME/bin:$PATH")
		os.Setenv("CLASSPATH", ".:$JAVA_HOME/lib/dt.jar:$JAVA_HOME/lib/tools.jar:")
		//go环境 go1.19.3
		os.Setenv("GOROOT", "env/go/go_linux")
		//windows
	} else if sysType == "windows" {
		//java环境 jdk1.8_352
		os.Setenv("JAVA_HOME", "env/jvm/jre_win")
		os.Setenv("PATH", "%JAVA_HOME%/bin;%PATH%")
		os.Setenv("CLASSPATH", ".;%JAVA_HOME%/lib/dt.jar;%JAVA_HOME%/lib/tools.jar;")
		//go环境 go1.19.3
		os.Setenv("GOROOT", "env/go/go_win")
	} else {
		logger.Log().Error(sysType, "系统不支持")
	}
}

func main() {
	a := app.NewWithID("com.mzydz.jarencryption")
	win = a.NewWindow("jar包加密")
	win.Resize(fyne.NewSize(800, 600))
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
