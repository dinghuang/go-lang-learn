package main

import (
	"fmt"
)

func main() {
	//结构体类似于其他语言中的 class，可以在结构体中定义多个字段，为结构体实现方法，实例化等。接下来我们定义一个结构体 Student，并为 Student 添加 name，age 字段，并实现 hello() 方法。
	//使用 Student{field: value, ...}的形式创建 Student 的实例，字段不需要每个都赋值，没有显性赋值的变量将被赋予默认值，例如 age 将被赋予默认值 0。
	stu := &Student{
		name: "Tom",
	}
	msg := stu.hello("Jack")
	// hello Jack, I am Tom
	fmt.Println(msg)
	//除此之外，还可以使用 new 实例化：
	stu2 := new(Student)
	// hello Alice, I am  , name 被赋予默认值""
	fmt.Println(stu2.hello("Alice"))
}

func (stu *Student) hello(person string) string {
	return fmt.Sprintf("hello %s, I am %s", person, stu.name)
}

type Student struct {
	name string
	age  int
}
