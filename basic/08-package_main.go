package main


import "fmt"

func main() {
	// 运行 go run main.go，会报错，add 未定义：
	//因为 go run main.go 仅编译 main.go 一个文件，所以命令需要换成
	//go run main.go calc.go  或者 go run .
	fmt.Println(addInPackage(3, 5))
	//Go 语言也有 Public 和 Private 的概念，粒度是包。如果类型/接口/方法/函数/字段的首字母大写，
	//则是 Public 的，对其他 package 可见，如果首字母小写，则是 Private 的，对其他 package 不可见。
}