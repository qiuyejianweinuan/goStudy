package main

import (
	"fmt"
)

// 关闭 一个通道意味着不能再向这个通道发送值了。
// 该特性可以向通道的接收方传达工作已经完成的信息
func main() {
	// 在这个例子中，我们将使用一个 jobs 通道，将工作内容。
	//从 main() 协程传递到一个工作协程中。 当我们没有更多的任务传递给工作协程时，我们将 close 这个 jobs 通道
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done
}
