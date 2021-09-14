package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	//=============参数与返回值===============
	quo, rem := div(100, 17)
	// 5 15
	fmt.Println(quo, rem)
	// 117
	fmt.Println(add(100, 17))
	//=============错误处理===============
	_, err := os.Open("filename.txt")
	if err != nil {
		fmt.Println(err)
	}
	//可以通过 errorw.New 返回自定义的错误
	if err := hello(""); err != nil {
		fmt.Println(err)
	}
	//error 往往是能预知的错误，但是也可能出现一些不可预知的错误，例如数组越界，这种错误可能会导致程序非正常退出，在 Go 语言中称之为 panic。
	fmt.Println(get(5))
	fmt.Println("finished")
	//在 Python、Java 等语言中有 try...catch 机制，在 try 中捕获各种类型的异常，在 catch 中定义异常处理的行为。Go 语言也提供了类似的机制 defer 和 recover。
	fmt.Println(getTry(5))
	fmt.Println("finished")
}

func getTry(index int) (ret int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Some error happened!", r)
			ret = -1
		}
	}()
	arr := [3]int{2, 3, 4}
	return arr[index]
}

func get(index int) int {
	arr := [3]int{2, 3, 4}
	return arr[index]
}

func hello(name string) error {
	if len(name) == 0 {
		return errors.New("error: name is null")
	}
	fmt.Println("Hello,", name)
	return nil
}

//实现2个数的加法（一个返回值）和除法（多个返回值）
func add(num1 int, num2 int) int {
	return num1 + num2
}

func div(num1 int, num2 int) (int, int) {
	return num1 / num2, num1 % num2
}

//也可以给返回值命名，简化 return，例如 add 函数可以改写为
func add2(num1 int, num2 int) (ans int) {
	ans = num1 + num2
	return
}
