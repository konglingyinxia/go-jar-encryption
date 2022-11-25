package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/konglingyinxia/go-jar-encryption/logger"
	"github.com/konglingyinxia/go-jar-encryption/resource"
)

var win fyne.Window

func main() {
	a := app.NewWithID("com.mzydz.jarencryption")
	win = a.NewWindow("jar包加密")
	//  //定义加密密码
	//定义文件选择框  //定义文件转码输出框
	layoutCreate()

	win.SetIcon(resource.LoadResource(resource.IconPath))
	win.SetMaster()
	win.Resize(fyne.NewSize(600, 450))
	logger.Log().Info("启动完成.............")
	win.ShowAndRun()
	logger.Log().Info("服务退出.............")
}

// 创建布局
func layoutCreate() {

	//自定义密码
	input := widget.NewEntry()
	input.SetText("20085151")
	line := container.NewVBox(widget.NewLabel("密码:"), input)
	win.SetContent(line)

}
