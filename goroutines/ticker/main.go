package main

import (
	"fmt"
	"time"
)

// 定时器是未来某一时刻执行一次的打点器
// 类似于js的setTimeout?

// setTimeout 允许我们将函数推迟到一段时间间隔之后再执行。
//setInterval 允许我们重复运行一个函数，从一段时间间隔之后开始运行，之后以该时间间隔连续重复运行该函数。

// - 打点器 则是为你想要以固定的时间间隔重复执行而准备的。 这里是一个打点器的例子，它将定时的执行，直到我们将它停止。

func main() {
	// 打点器和定时器的机制有点相似：使用一个通道来发送数据。
	//这里我们使用通道内建的 select，等待每 500ms 到达一次的值。
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)
	// 启动协程异步等待接受 t

	for _, i := range []int{1, 2, 4, 5} {
		if i == 2 {
			break
		}
		fmt.Println("i", i)
	}
	go func() {
		for {
			select {
			//当done为true时,退出
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1400 * time.Millisecond)
	// 停止打点器打点
	ticker.Stop()
	//传递
	done <- true
	fmt.Println("ticker stopped")
}
