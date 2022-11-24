package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"go-jar-encryption/resource"
)

func main() {
	a := app.NewWithID("com.mzydz.jarencryption")
	win := a.NewWindow("jar包加密")

	//  //定义加密密码
	//定义文件选择框  //定义文件转码输出框

	win.SetIcon(resource.LoadResource(resource.IconPath))
	win.Resize(fyne.NewSize(600, 450))
	win.ShowAndRun()
}
