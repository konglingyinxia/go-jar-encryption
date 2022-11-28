# 1、说明

本项目为springboot项目 jar 自动加密项目,可防止项目发出的项目代码反编译，从而泄露源码

# 2、功能特点

- 1、对`xjar` 项目进行封装
- 2、使用golang编写pc客户端，方便运维操作，
- 3、操作界面丰富，上手简单，可编译出`win`,`linux` 平台的可执行加密后的jar包

# 打包

- 1、准备静态资源
    - 1、使用`go-bindata`打包静态资源文件
      ```aidl
      安装：go get github.com/jteeuwen/go-bindata/...
      执行：go-bindata -o=./asset/asset.go -pkg=asset config/...
 
      ```
    - 2、交叉编译
        - linux
          ```
             CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o abc-demo-linux main.go
             CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o adc-demo-mac main.go
             CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o abc-demo-windows.exe main.go
          ```
        - win
           ```
             Windows主机编译Windows客户端

             SET CGO_ENABLED=0
             SET GOOS=windows
             SET GOARCH=amd64
             go build -o abc-demo-windows.exe main.go
                                             
             Windows主机编译LINUX客户端
                                             
             SET CGO_ENABLED=0
             SET GOOS=linux
             SET GOARCH=amd64
             go build -o abc-demo-linux main.go

           ```
