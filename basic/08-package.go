package main



func addInPackage(num1 int, num2 int) int {
	//一般来说，一个文件夹可以作为 package，同一个 package 内部变量、类型、方法等定义可以相互看到。
	//比如我们新建一个文件 calc.go， main.go 平级，分别定义 add 和 main 方法。
	return num1 + num2
}
