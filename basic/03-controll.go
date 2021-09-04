package main

import (
	"fmt"
)

func main() {
	//条件语句 if else
	age := 18
	if age < 18 {
		fmt.Printf("Kid")
	} else {
		fmt.Printf("Adult")
	}

	// 可以简写为：
	if age := 18; age < 18 {
		fmt.Printf("Kid")
	} else {
		fmt.Printf("Adult")
	}

	//switch
	//在这里，使用了type 关键字定义了一个新的类型 Gender。
	type Gender int8
	//使用 const 定义了 MALE 和 FEMALE 2 个常量，Go 语言中没有枚举(enum)的概念，一般可以用常量的方式来模拟枚举。
	const (
		MALE   Gender = 1
		FEMALE Gender = 2
	)

	gender := MALE

	switch gender {
	case FEMALE:
		fmt.Println("female")
	case MALE:
		fmt.Println("male")
	default:
		fmt.Println("unknown")
	}
	//和其他语言不同的地方在于，Go 语言的 switch 不需要 break，匹配到某个 case，执行完该 case 定义的行为后，默认不会继续往下执行。如果需要继续往下执行，需要使用 fallthrough，例如：
	switch gender {
	case FEMALE:
		fmt.Println("female")
		fallthrough
	case MALE:
		fmt.Println("male")
		fallthrough
	default:
		fmt.Println("unknown")
	}
	// 输出结果
	// male
	// unknown

	//for 循环
	//一个简单的累加的例子，break 和 continue 的用法与其他语言没有区别。
	sum := 0
	for i := 0; i < 10; i++ {
		if sum > 50 {
			break
		}
		sum += i
	}
	//对数组(arr)、切片(slice)、字典(map) 使用 for range 遍历：
	nums := []int{10, 20, 30, 40}
	for i, num := range nums {
		fmt.Println(i, num)
	}
	// 0 10
	// 1 20
	// 2 30
	// 3 40
	m2 := map[string]string{
		"Sam":   "Male",
		"Alice": "Female",
	}

	for key, value := range m2 {
		fmt.Println(key, value)
	}
	// Sam Male
	// Alice Female

}
