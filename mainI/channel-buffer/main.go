package main

import (
	"fmt"
)

// 默认情况下，通道是 无缓冲 的，这意味着只有对应的接收（<- chan）
// 通道准备好接收时，才允许进行发送（chan <-）。

// 有缓冲通道 允许在没有对应接收者的情况下，缓存一定数量的值
func main() {

	fmt.Println()
	// 这里我们 make 了一个字符串通道，最多允许缓存 2 个值
	messages := make(chan string, 2)

	// 这里我们 make 了一个字符串通道，最多允许缓存 2 个值
	messages <- "buffered"
	fmt.Println(<-messages)
	messages <- "channel"
	messages <- "channel1"

	// 然后我们可以正常接收这两个值。
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
