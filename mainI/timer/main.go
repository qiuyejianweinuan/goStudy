package main

import (
	"fmt"
	"time"
)

// 定时器
func main() {
	// 定时器表示在未来某一时刻的独立事件。
	//你告诉定时器需要等待的时间,然后他将提供一个用于通知的通道.这里的定时器将等待两秒
	timer1 := time.NewTimer(2 * time.Second)

	// timer1.C会一直阻塞,直到定时器的通道C明确的发送了定时器失效的值
	<-timer1.C
	//打印提示信息
	fmt.Println("Timer 1 fired ")

	//声明timer2
	timer2 := time.NewTimer(time.Second)

	// 如果单纯等待,使用time.Sleep就可以,使用定时器的原因就是定时器在触发之前可以提前取消
	// 下面是启动一个go协程用来等待timer2的触发,然后使用stop方法提前将其取消
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	//测试阻止timer2被停止
	time.Sleep(2 * time.Second)
	//取消timer2

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(2 * time.Second)
}
