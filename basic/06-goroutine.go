package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 创建大小为 10 的缓冲信道
var ch = make(chan string, 10)

func main() {
	//Go 语言提供了 sync 和 channel 两种方式支持协程(goroutine)的并发。
	//=====================================sync==================================
	//例如希望并发下载 N 个资源，多个并发协程之间不需要通信，那么就可以使用 sync.WaitGroup，等待所有并发协程执行结束。
	for i := 0; i < 3; i++ {
		//为 wg 添加一个计数，wg.Done()，减去一个计数。
		wg.Add(1)
		//启动新的协程并发执行 download 函数。
		go download("a.com/" + string(i+'0'))
	}

	//等待所有的协程执行结束。
	wg.Wait()
	fmt.Println("Done!")
	//=====================================channel==================================
	//使用 channel 信道，可以在协程之间传递消息。阻塞等待并发协程返回消息。
	for i := 0; i < 3; i++ {
		//为 wg 添加一个计数，wg.Done()，减去一个计数。
		wg.Add(1)
		//启动新的协程并发执行 download 函数。
		go downloadChannel("a.com/" + string(i+'0'))
	}
	for i := 0; i < 3; i++ {
		msg := <-ch // 等待信道返回消息。
		fmt.Println("finish", msg)
	}
	fmt.Println("Done!")
}

func downloadChannel(url string) {
	fmt.Println("start to download", url)
	time.Sleep(time.Second)
	// 将 url 发送给信道
	ch <- url
}

func download(url string) {
	fmt.Println("start to download", url)
	// 模拟耗时操作
	time.Sleep(time.Second)
	wg.Done()
}
