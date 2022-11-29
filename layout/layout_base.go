package layout

import (
	"bufio"
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/konglingyinxia/go-jar-encryption/logger"
	"github.com/konglingyinxia/go-jar-encryption/projectpath"
	"github.com/ncruces/zenity"
	"io"
	"io/ioutil"
	os2 "os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var (
	javaBin = ""
	goBin   = ""
)

func init() {
	//初始化系统环境变量
	sysType := runtime.GOOS
	if sysType == "linux" {
		javaBin = "env/jvm/jre_linux/bin/java"
		goBin = "env/go/go_linux/bin/go"
		//windows
	} else if sysType == "windows" {
		javaBin = "env/jvm/jre_win/bin/java.exe"
		goBin = "env/go/go_win/bin/go.exe"
	} else {
		logger.Log().Error(sysType, "系统不支持")
	}
}

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

// BaseFrom  创建布局 	//  //定义加密密码
//
//	//定义文件选择框  //定义文件转码输出框
func BaseFrom(win fyne.Window) {
	//自定义密码
	input := widget.NewEntry()
	input.Validator = validation.NewRegexp("^[A-Za-z0-9]{6,8}$", "只能包含字母、数字，长度6到8")
	outFileName := widget.NewEntry()
	outPathInput := widget.NewEntry()
	openFileInput := widget.NewEntry()
	openFile := widget.NewButton("选择", func() {
		jarFile, err := zenity.SelectFile(
			zenity.FileFilters{
				{"jar files", []string{"*.jar"}},
			},
		)
		if err == zenity.ErrCanceled {
			logger.Log().Error("用户取消了jar包选择", err)
			zenity.Info("您取消了", zenity.Title("Information"))
			return
		}
		if err != nil {
			logger.Log().Error("jar包选择错误", err)
			zenity.Info(err.Error())
			return
		}
		logger.Log().Info("原始JAR包地址：", jarFile)
		openFileInput.Bind(binding.BindString(&jarFile))
	})
	outPath := widget.NewButton("选择", func() {
		outPath, err := zenity.SelectFile(
			zenity.Directory(),
		)
		if err == zenity.ErrCanceled {
			logger.Log().Error("用户取消了输出目录选择", err)
			zenity.Info("您取消了输出目录选择")
			return
		}
		if err != nil {
			logger.Log().Error("输出目录选择错误", err)
			zenity.Error("输出目录选择错误")
			return
		}
		if !DirEmpty(outPath) {
			logger.Log().Error("请选择空目录", err)
			//dialog.ShowError(err, win)
			zenity.Error("请选择空目录")
			return
		}
		logger.Log().Info("加密JAR包输出目录地址：", outPath)
		outPathInput.Bind(binding.BindString(&outPath))
		fileName := filepath.Base(openFileInput.Text)
		outFileName.Bind(binding.BindString(&fileName))
	})

	inItem := container.NewHSplit(openFileInput, openFile)
	inItem.SetOffset(0.85)
	outItem := container.NewHSplit(outPathInput, outPath)
	outItem.SetOffset(0.85)
	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "密码", Widget: input, HintText: "输入加密密码"},
			{Text: "系统类型", Widget: osType, HintText: "选择系统类型"},
			{Text: "jar包选择", Widget: inItem, HintText: "待加密的jar包地址"},
			{Text: "输出目录", Widget: outItem, HintText: "输出目录"},
			{Text: "文件名", Widget: outFileName, HintText: "输出文件名"},
		},
		OnSubmit: func() {
			Confirm(win, "确定提交", func(b bool) {
				if b {
					logger.Log().Info("你提交了加密参数...")
					pwd := input.Text
					os := osType.Text
					inputJar := openFileInput.Text
					outP := outPathInput.Text
					outFileNameP := outFileName.Text
					if pwd == "" || os == "" || outP == "" || inputJar == "" || outFileNameP == "" {
						ShowError(errors.New("输入框都不能为空"), win)
					} else {
						s := os2.PathSeparator
						outP = outP + string(s) + outFileNameP
						logger.Log().Info("平台：", os, ",原始：", inputJar, "，输出：", outP, ",加密开始...")
						err := encodeBuild(pwd, os, inputJar, outP, win)
						if err != nil {
							logger.Log().Info("平台：", os, ",原始：", inputJar, "，输出：", outP, ",加密失败...")
							ShowError(err, win)
							return
						}
						logger.Log().Info("平台：", os, ",原始：", inputJar, "，输出：", outP, ",加密成功...")
					}
				} else {
					logger.Log().Info("二次确认-您取消了jar包加密")
				}
			})
		},
		SubmitText: "确认",
	}
	richText := widget.NewMultiLineEntry()
	richText.Wrapping = fyne.TextWrapWord
	go ReadLog(richText)
	label := widget.NewLabel("日志")
	vHBox := container.NewVBox(label, widget.NewSeparator(), container.New(layout.NewGridWrapLayout(fyne.Size{Width: 800, Height: 350}), richText))
	b := container.NewVBox(form, widget.NewSeparator(), vHBox)
	Base(win, b)
}

func encodeBuild(pwd string, os string, jarFile string, outFileName string, win fyne.Window) error {
	outDir := filepath.Dir(outFileName)
	//执行jar加密
	err := jarEncode(pwd, jarFile, outFileName)
	if err != nil {
		logger.Log().Error("jar包加密失败....", err)
		ShowError(errors.New("jar包加密失败"), win)
		return err
	}
	xjarGoPath := filepath.Join(outDir, "xjar.go")
	logger.Log().Info(xjarGoPath, "编译开始.......")
	//打包成linux
	if os == linuxOx {
		buildLinux(xjarGoPath)
		//打包成win执行包
	} else if os == winOs {
		buildWin(xjarGoPath)
	} else if os == allOs {
		buildLinux(xjarGoPath)
		buildWin(xjarGoPath)
	} else {
		return errors.New("不支持的系统类型")
	}
	logger.Log().Info(xjarGoPath, "编译结束......")
	return nil

}

// GOARCH=amd64  GOOS=linux  go build  xjar.go
func buildLinux(xjarGoPath string) {
	log := make(chan string)
	dir := filepath.Dir(xjarGoPath)
	os2.Setenv("GOARCH", "amd64")
	os2.Setenv("GOOS", "linux")
	cmdXjarGo := exec.Command("go", "build", xjarGoPath)
	cmdXjarGo.Dir = dir
	go cmdExec(cmdXjarGo, log)
	for {
		str := <-log
		if str == "&end|end|end&" {
			break
		} else {
			logger.Log().Info(str)
		}
	}
}

// `GOARCH=amd64  GOOS=windows  go build  xjar.go`
func buildWin(xjarGoPath string) {
	log := make(chan string)
	dir := filepath.Dir(xjarGoPath)
	os2.Setenv("GOARCH", "amd64")
	os2.Setenv("GOOS", "windows")
	cmdXjarGo := exec.Command(projectpath.RootPath()+"/"+goBin, "build", xjarGoPath)
	cmdXjarGo.Dir = dir
	go cmdExec(cmdXjarGo, log)
	for {
		str := <-log
		if str == "&end|end|end&" {
			break
		} else {
			logger.Log().Info(str)
		}
	}

}

// 参数顺序为：filePath=? pwd=?  outPath=?
func jarEncode(pwd string, file string, outPath string) error {
	cmdJava := exec.Command(projectpath.RootPath()+"/"+javaBin, "-jar", "lib/tools-jar.jar", "filePath="+file,
		"pwd="+pwd, "outPath="+outPath)
	log := make(chan string)
	go cmdExec(cmdJava, log)
	logger.Log().Info("jar包加密日志开始..........................")
	for {
		str := <-log
		if str == "&end|end|end&" {
			logger.Log().Info("jar包加密日志结束..........................")
			break
		} else {
			logger.Log().Info(str)
		}
	}
	//判断文件是否生成
	if !PathExists(outPath) {
		return errors.New(outPath + ",加密文件未正常生成")
	}
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

func PathExists(path string) bool {
	_, err := os2.Stat(path)
	if err == nil {
		return true
	}
	if os2.IsNotExist(err) == false {
		return true
	}
	return false
}

// DirEmpty 空目录：true ,非空目录：false
func DirEmpty(dirPath string) bool {
	dir, _ := ioutil.ReadDir(dirPath)
	if len(dir) == 0 {
		return true
	} else {
		return false
	}
}

func ReadLog(logRich *widget.Entry) {
	for {
		path := logger.LogFile()
		file, err := os2.Open(path)
		if err != nil {
			fmt.Println("打开文件出错：", err)
		}
		file.Seek(0, io.SeekEnd)
		reader := bufio.NewReader(file)
		for {
			line, _, err := reader.ReadLine()
			if err == io.EOF {
				time.Sleep(time.Second)
			} else {
				str := logRich.Text + "\n" + string(line)
				logRich.Bind(binding.BindString(&str))
				logRich.Refresh()
			}
		}
	}

}
