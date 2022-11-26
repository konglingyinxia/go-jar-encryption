package main

import (
	"bufio"
	"fyne.io/fyne/v2"
	"io"
	"log"
	"os/exec"
	"strings"
)

var win fyne.Window

func init() {
	//初始化系统环境变量
	//java环境
	//go环境
}

//func main() {
//	a := app.NewWithID("com.mzydz.jarencryption")
//	win = a.NewWindow("jar包加密")
//	win.Resize(fyne.NewSize(600, 450))
//	win.SetIcon(resource.LoadResource(resource.IconPath))
//
//	layout.BaseCreate(win)
//
//	logger.Log().Info("启动完成.............")
//	win.SetMaster()
//	// 窗口居中
//	win.CenterOnScreen()
//	win.ShowAndRun()
//	logger.Log().Info("服务退出.............")
//}

func main() {
	cmdJava := exec.Command("java", "-jar", "lib/tools-test.jar")
	stdout, _ := cmdJava.StdoutPipe()
	cmdJava.Start()
	reader := bufio.NewReader(stdout)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil || io.EOF == err {
			break
		}
		log.Println(line)
	}
	cmdJava.Wait()
}
