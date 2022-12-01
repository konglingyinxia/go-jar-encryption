package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
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
	//start()
	myApp := app.New()
	myWindow := myApp.NewWindow("List Data")

	data := binding.BindStringList(
		&[]string{"Item 1", "Item 2", "Item 3"},
	)

	list := widget.NewListWithData(data,
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*widget.Label).Bind(i.(binding.String))
		})

	add := widget.NewButton("Append", func() {
		val := fmt.Sprintf("Item %d", data.Length()+1)
		data.Append(val)
	})
	myWindow.SetContent(container.NewBorder(nil, add, nil, nil, list))
	myWindow.ShowAndRun()
}
