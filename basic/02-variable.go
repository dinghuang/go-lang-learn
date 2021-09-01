package main

import (
	"fmt"
	"reflect"
)

func main() {
	//===============变量===============

	// 如果没有赋值，默认为0
	var a int
	fmt.Println(a)
	// 声明时赋值
	var b int = 1
	fmt.Println(b)
	// 声明时赋值
	var c = 1
	fmt.Println(c)
	//var d = 1，因为 1 是 int 类型的，所以赋值时，a 自动被确定为 int 类型，所以类型名可以省略不写，这种方式还有一种更简单的表达：
	d := 1
	fmt.Println(d)
	var 速度 int
	// keep the compiler happy
	_ = 速度

	//===============简单类型===============

	/**
	空值：nil
	整型类型： int(取决于操作系统), int8, int16, int32, int64, uint8, uint16, …
	浮点数类型：float32, float64
	字节类型：byte (等价于uint8)
	字符串类型：string
	布尔值类型：boolean，(true 或 false)
	**/
	var a1 int8 = 10
	fmt.Println(a1)
	var c1 byte = 'a'
	fmt.Println(c1)
	var b1 float32 = 12.2
	fmt.Println(b1)
	var msg = "Hello World"
	fmt.Println(msg)
	ok := false
	//基础类型
	fmt.Println(ok)
	fmt.Println(42, 8500, 344433, -2323)
	fmt.Println(3.14, 6.28, -42.)
	fmt.Println(true, false)
	fmt.Println("Hi! I'm Inanc!")
	fmt.Println("Merhaba, adım İnanç!")
	fmt.Println(0x0, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9)
	fmt.Println(0xa, 0xb, 0xc, 0xd, 0xe, 0xf)
	// 17
	fmt.Println(0x11)
	// 25
	fmt.Println(0x19)
	// 50
	fmt.Println(0x32)
	// 100
	fmt.Println(0x64)
	//===============字符串===============
	// 在 Go 语言中，字符串使用 UTF8 编码，UTF8 的好处在于，如果基本是英文，每个字符占 1 byte，和 ASCII 编码是一样的，非常节省空间，如果是中文，一般占3字节。包含中文的字符串的处理方式与纯 ASCII 码构成的字符串有点区别。
	str1 := "Golang"
	str2 := "Go语言"
	//reflect.TypeOf().Kind() 可以知道某个变量的类型，我们可以看到，字符串是以 byte 数组形式保存的，类型是 uint8，占1个 byte，打印时需要用 string 进行类型转换，否则打印的是编码值。
	// uint8
	fmt.Println(reflect.TypeOf(str2[2]).Kind())
	//因为字符串是以 byte 数组的形式存储的，所以，str2[2] 的值并不等于语。str2 的长度 len(str2) 也不是 4，而是 8（ Go 占 2 byte，语言占 6 byte）。
	// 108 l
	fmt.Println(str1[2], string(str1[2]))
	// 232 è
	fmt.Printf("%d %c\n", str2[2], str2[2])
	// len(str2)： 8
	fmt.Println("len(str2)：", len(str2))
	//正确的处理方式是将 string 转为 rune 数组
	runeArr := []rune(str2)
	// int32
	fmt.Println(reflect.TypeOf(runeArr[2]).Kind())
	// 35821 语
	fmt.Println(runeArr[2], string(runeArr[2]))
	// len(runeArr)： 4
	fmt.Println("len(runeArr)：", len(runeArr))

	//===============变量===============

}
