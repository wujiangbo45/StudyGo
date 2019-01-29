package main

import (
	"fmt"
	"time"
)

/**
模拟心跳和超时
每隔2秒发送一次心跳,发送5次之后停止,然后等待10秒后发出超时
*/
func main() {
	mch := make(chan string)
	// 心跳发送
	go func(ch chan string) {
		for i := 0; i < 5; i++ {
			mch <- "ping"
			time.Sleep(2 * time.Second)
		}
	}(mch)

	// 处理通道信息
	go func(ch chan string) {
		for {
			select {
			case m := <-ch:
				fmt.Println(m)
			case <-time.After(10 * time.Second):
				fmt.Println("timeout")
			}
		}

	}(mch)
	time.Sleep(500 * time.Second)
}
