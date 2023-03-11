package main

import "fmt"

// 通道本身没有方向,我们可以根据自定义函数来控制通道是接受我们的数据还是我们从通道中读取数据

func ping(pings chan<- string, msg string) {
	// 此函数的 pings 我们向他发送数据的,是一个只写通道
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	// 此函数中的 pings ：我们从pings中接收数据并没有向他写入数据,是一个只读通道
	msg := <-pings
	// 我们向 pongs
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
