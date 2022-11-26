package layout

import (
	"bufio"
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/konglingyinxia/go-jar-encryption/logger"
	"io"
	"os/exec"
	"strings"
)

func Confirm(win fyne.Window, msg string, callback func(bool)) {
	cnf := dialog.NewConfirm("再次确认", msg, callback, win)
	cnf.SetDismissText("取消")
	cnf.SetConfirmText("确定")
	cnf.Show()
}

// ShowError shows a dialog over the specified window for an application error.
// The message is extracted from the provided error (should not be nil).
func ShowError(err error, parent fyne.Window) {
	d := dialog.NewError(err, parent)
	d.Show()
}

const linuxOx = "linux"
const winOs = "win"
const allOs = "all"

var osType = widget.NewSelectEntry([]string{linuxOx, winOs, allOs})

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
					pwd := input.Text
					os := osType.Text
					inputJar := openFileInput.Text
					outP := outPathInput.Text
					if pwd == "" || os == "" || outP == "" || inputJar == "" {
						ShowError(errors.New("输入框都不能为空"), win)
					} else {
						logger.Log().Info("平台：", os, ",原始：", inputJar, "，输出：", outP, ",加密开始...")
						encodeBuild(pwd, os, inputJar, outP, win)
						logger.Log().Info("平台：", os, ",原始：", inputJar, "，输出：", outP, ",加密结束...")
					}
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

func encodeBuild(pwd string, os string, jarFile string, out string, win fyne.Window) {
	//执行jar加密
	err := jarEncode(pwd, jarFile, out, win)
	if err != nil {
		logger.Log().Error("jar包加密失败....")
		return
	}

	//打包成linux
	if os == linuxOx {
		//打包成win执行包
	} else if os == winOs {

	} else if os == allOs {

	} else {
		ShowError(errors.New("系统类型选择错误"), win)
	}

}

// 参数顺序为：filePath=? pwd=?  outPath=?
func jarEncode(pwd string, file string, outPath string, win fyne.Window) error {
	cmdJava := exec.Command("java", "-jar", "lib/tools-jar.jar", file, pwd, outPath)
	log := make(chan string)
	go cmdExec(cmdJava, log)
	go func() {
		logger.Log().Info("jar包加密日志开始..........................")
		for {
			str := <-log
			if str == "&end|end|end&" {
				logger.Log().Info("jar包加密日志结束..........................")
			} else {
				logger.Log().Info(str)
			}

		}
	}()
	return nil
}

func cmdExec(cmd *exec.Cmd, chanel chan string) {
	stdout, _ := cmd.StdoutPipe()
	cmd.Start()
	reader := bufio.NewReader(stdout)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil || io.EOF == err {
			chanel <- "&end|end|end&"
			break
		}
		chanel <- line
	}
}
