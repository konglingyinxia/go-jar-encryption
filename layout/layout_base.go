package layout

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/konglingyinxia/go-jar-encryption/logger"
)

func Confirm(win fyne.Window, msg string, callback func(bool)) {
	cnf := dialog.NewConfirm("再次确认", msg, callback, win)
	cnf.SetDismissText("取消")
	cnf.SetConfirmText("确定")
	cnf.Show()
}

const linuxOx = "linux"
const winOs = "win"

func Base(win fyne.Window, c fyne.CanvasObject) {
	a := fyne.CurrentApp()
	themes := container.NewGridWithColumns(2,
		widget.NewButton("Dark", func() {
			a.Settings().SetTheme(theme.DarkTheme())
		}),
		widget.NewButton("Light", func() {
			a.Settings().SetTheme(theme.LightTheme())
		}),
	)
	r := container.NewBorder(nil, themes, nil, nil, c)
	win.SetContent(r)
}

// BaseCreate  创建布局 	//  //定义加密密码
//
//	//定义文件选择框  //定义文件转码输出框
func BaseCreate(win fyne.Window) {
	//自定义密码
	input := widget.NewEntry()
	input.Validator = validation.NewRegexp("^[A-Za-z0-9]{6,8}$", "只能包含字母、数字，长度6到8")
	osType := widget.NewSelectEntry([]string{linuxOx, winOs})
	openFileInput := widget.NewEntry()
	openFile := widget.NewButton("选择", func() {
		fd := dialog.NewFileOpen(func(read fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, win)
				return
			}
			if read == nil {
				logger.Log().Info("用户取消原始jar包选择")
			}
			if read != nil {
				path := read.URI().Path()
				logger.Log().Info("原始JAR包地址：", path)
				openFileInput.Bind(binding.BindString(&path))
			}
		}, win)
		fd.SetFilter(storage.NewExtensionFileFilter([]string{".jar"}))
		fd.Show()
	})
	outPathInput := widget.NewEntry()
	outPath := widget.NewButton("选择", func() {
		fd := dialog.NewFolderOpen(func(read fyne.ListableURI, err error) {
			if err != nil {
				dialog.ShowError(err, win)
				return
			}
			if read == nil {
				logger.Log().Info("用户取消输出目录选择")
			}
			if read != nil {
				path := read.Path()
				logger.Log().Info("加密JAR包输出目录地址：", path)
				outPathInput.Bind(binding.BindString(&path))
			}
		}, win)
		fd.Show()
	})
	inItem := container.NewHSplit(openFileInput, openFile)
	inItem.SetOffset(0.85)
	outItem := container.NewHSplit(outPathInput, outPath)
	outItem.SetOffset(0.85)
	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "密码", Widget: input, HintText: "输入加密密码"},
			{Text: "系统类型", Widget: osType, HintText: "选择系统类型"},
			{Text: "文件目录", Widget: inItem, HintText: "待加密的jar包地址"},
			{Text: "输出目录", Widget: outItem, HintText: "输出目录"},
		},
		OnSubmit: func() {
			Confirm(win, "确定提交", func(b bool) {
				if b {
					logger.Log().Info("你提交了加密参数...")
					inFile := input.Text
					logger.Log().Info(inFile)
				} else {
					logger.Log().Info("二次确认您取消了jar包加密")
				}
			})
		},
		SubmitText: "确认",
	}
	form.Resize(fyne.NewSize(200, 200))
	b := container.NewVBox(form)
	Base(win, b)
}
