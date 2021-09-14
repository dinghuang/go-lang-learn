# go lang 学习

## 安装
本人是macOS，所以安装就按照macOS来安装。前提是已经安装了homeBrew
```
$ brew install go
```
查看go安装情况
```
$ go version
go version go1.15.2 darwin/amd64
```
从 Go 1.11 版本开始，Go 提供了 [Go Modules](https://github.com/golang/go/wiki/Modules) 的机制，推荐设置以下环境变量，第三方包的下载将通过国内镜像，避免出现官方网址被屏蔽的问题。
```
$ go env -w GOPROXY=https://goproxy.cn,direct
```
查看环境变量
```
$ go env
```

## IDE工具

如果用的VSCODE的话，文件夹命名别加``-``，安装go插件，安装Code Runner插件
## 语法学习

这部分包含了基础的go lang语法，在目录basic下面
### Hello World

- ``01-main`` Hello World程序
- ``02-variable`` 变量与内置数据类型
- ``03-controll`` 流程控制(if, for, switch)
- ``04-function`` 函数(functions)

