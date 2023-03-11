package main

import (
	"fmt"
	"sync"
	"time"
)

// 非阻塞通道操作
// 常规的通过通道发送和接收数据是阻塞的。
// 然而，我们可以使用带一个 default 子句的 select 来实现 非阻塞 的发送、接收
// 甚至是非阻塞的多路 select。
func main() {
	var wg sync.WaitGroup
	messages := make(chan string, 1)
	signals := make(chan bool)

	wg.Add(1)
	go func() {
		select {
		case msg := <-messages:
			fmt.Println("received message", msg)
		default:
			fmt.Println("no message received")
		}
		wg.Done()
	}()

	time.Sleep(1 * time.Second)
	messages <- "hi"
	fmt.Println(<-messages)
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}

	wg.Wait()
}
