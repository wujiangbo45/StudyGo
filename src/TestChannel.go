package main

import (
	"fmt"
)

/**
模拟python的生成器,实现累加器
*/
func xrange() chan int {
	var ch = make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i // i写入通道会阻塞
		}
	}()
	return ch
}

func main() {
	// 接收通道信息并阻塞下一条记录
	chout := xrange()
	fmt.Println(<-chout)

	ch1 := SendForChannel()
	ch1 <- "login"
	abc := <-ch1
	ch1 <- "login11"
	fmt.Println(abc)
	fmt.Println(<-ch1)
}

/**
模拟心跳发送ping返回pong
*/
func Notification(uname string) chan string {
	ch := make(chan string)
	go func() {
		ch <- fmt.Sprintf("pong back %s", uname)
	}()
	return ch
}

/**
用户名登录,返回通知channel
每一个send都会返回一个channel
*/
func Send(uname string) chan string {
	SendCh := make(chan string)
	go func(ch chan string) {
		recCh := Notification(uname)
		ch <- <-recCh
	}(SendCh)
	fmt.Println(uname + " ping!")
	return SendCh
}

func NotificationByChannel() chan string {
	// 通知频道
	NoCh := make(chan string)
	go func(ch chan string) {
		for {
			// 接收信息
			str := <-ch
			fmt.Printf("%s send ping package... \n", str)
			ch <- fmt.Sprintf("pong back %s", str)
			// 信息回写频道
		}
	}(NoCh)
	return NoCh
}

/**
 通过channel通信,通信均使用同一个channel
	ch1 := SendForChannel()
	ch1 <- "login"
	abc := <- ch1
	ch1 <- "login11"
	fmt.Println(abc)
	fmt.Println(<-ch1)
*/
func SendForChannel() chan string {
	ch := make(chan string)
	// 接收通知频道信息
	chre := NotificationByChannel()
	go func(chre chan string) {
		for {
			// 登录频道收到消息传递给通知频道
			chre <- <-ch
			// 接收通知频道的消息通知到登录频道
			ch <- <-chre
		}

	}(chre)
	return ch
}
