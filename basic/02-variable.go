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

	//===============数组(array)与切片(slice)===============
	//声明数组
	// 一维
	var arr [5]int
	// 二维
	var arr2 [5][5]int
	_ = arr
	_ = arr2
	//声明时初始化
	arr = [5]int{1, 2, 3, 4, 5}
	//使用 [] 索引/修改数组
	for i := 0; i < len(arr); i++ {
		arr[i] += 100
	}
	fmt.Println(arr)
	//数组的长度不能改变，如果想拼接2个数组，或是获取子数组，需要使用切片。切片是数组的抽象。 切片使用数组作为底层结构。切片包含三个组件：容量，长度和指向底层数组的指针,切片可以随时进行扩展
	//声明切片：
	slice1 := make([]float32, 0) // 长度为0的切片
	_ = slice1
	slice2 := make([]float32, 3, 5)       // [0 0 0] 长度为3容量为5的切片
	fmt.Println(len(slice2), cap(slice2)) // 3 5
	//使用切片
	// 添加元素，切片容量可以根据需要自动扩展
	slice2 = append(slice2, 1, 2, 3, 4)   // [0, 0, 0, 1, 2, 3, 4]
	fmt.Println(len(slice2), cap(slice2)) // 7 12
	// 子切片 [start, end)
	sub1 := slice2[3:]  // [1 2 3 4]
	sub2 := slice2[:3]  // [0 0 0]
	sub3 := slice2[1:4] // [0 0 1]
	// 合并切片
	//声明切片时可以为切片设置容量大小，为切片预分配空间。在实际使用的过程中，如果容量不够，切片容量会自动扩展。
	//sub2... 是切片解构的写法，将切片解构为 N 个独立的元素。

	combined := append(sub1, sub2...) // [1, 2, 3, 4, 0, 0, 0]
	_ = sub3
	_ = combined

	//=============== 字典(键值对，map)===============
	//map 类似于 java 的 HashMap，Python的字典(dict)，是一种存储键值对(Key-Value)的数据解构。使用方式和其他语言几乎没有区别。
	// 仅声明
	m1 := make(map[string]int)
	// 声明时初始化
	m2 := map[string]string{
		"Sam":   "Male",
		"Alice": "Female",
	}
	_ = m2
	// 赋值/修改
	m1["Tom"] = 18

	//=============== 指针(pointer)===============
	//指针即某个值的地址，类型定义时使用符号*，对一个已经存在的变量，使用 & 获取该变量的地址。
	str := "Golang"
	var p *string = &str // p 是指向 str 的指针
	*p = "Hello"
	fmt.Println(str) // Hello 修改了 p，str 的值也发生了改变
	//一般来说，指针通常在函数传递参数，或者给某个类型定义新的方法时使用。Go 语言中，参数是按值传递的，如果不使用指针，函数内部将会拷贝一份参数的副本，对参数的修改并不会影响到外部变量的值。如果参数使用指针，对参数的传递将会影响到外部变量。
	num := 100
	add(num)
	fmt.Println(num) // 100，num 没有变化

	realAdd(&num)
	fmt.Println(num) // 101，指针传递，num 被修改
}

func add(num int) {
	num += 1
}

func realAdd(num *int) {
	*num += 1
}
