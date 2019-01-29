package main

import (
	"fmt"
	"time"
)

/**
goroutine简单使用以及goroutine间通信
*/
func main() {
	// 定义接收str类型的channel
	ch := make(chan string)
	go func() {
		// test 写入channel
		ch <- "test"
	}()
	go func() {
		// 接收channel信息
		for i := range ch {
			fmt.Println(i)
		}

	}()
	ch <- "hello"
	// 睡眠等待,否则程序执行太快无法看到输出
	time.Sleep(15 * time.Second)

}
