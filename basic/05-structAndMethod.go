package main

import (
	"fmt"
)

func main() {
	//=====================================struct==================================
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

	//=====================================interfaces==================================
	//Go 语言中，并不需要显式地声明实现了哪一个接口，只需要直接实现该接口对应的方法即可。
	//实例化 Student后，强制类型转换为接口类型 Person。
	//将空值 nil 转换为 *Student 类型，再转换为 Person 接口，如果转换失败，说明 Student 并没有实现 Person 接口的所有方法。
	//Worker 同上。
	var p Person = &Student{
		name: "Tom",
		age:  18,
	}
	// Tom
	fmt.Println(p.getName())
	//实例可以强制类型转换为接口，接口也可以强制类型转换为实例。
	stuP := p.(*Student)
	fmt.Println(stuP.getAge())

	//空接口
	//如果定义了一个没有任何方法的空接口，那么这个接口可以表示任意类型。例如
	m := make(map[string]interface{})
	m["name"] = "Tom"
	m["age"] = 18
	m["scores"] = [3]int{98, 99, 85}
	// map[age:18 name:Tom scores:[98 99 85]]
	fmt.Println(m)
}

func (stu *Student) hello(person string) string {
	return fmt.Sprintf("hello %s, I am %s", person, stu.name)
}

type Student struct {
	name string
	age  int
}

type Person interface {
	getName() string
}

func (stu *Student) getName() string {
	return stu.name
}

func (stu *Student) getAge() int {
	return stu.age
}

type Worker struct {
	name   string
	gender string
}

func (w *Worker) getName() string {
	return w.name
}
