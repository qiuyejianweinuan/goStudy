package main

import (
	"fmt"
	"sync"
	"time"
)

// worker函数
func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	// 睡眠一秒钟，以此来模拟耗时的任务。
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	// 这个 WaitGroup 用于等待这里启动的所有协程完成。
	//注意：如果 WaitGroup 显式传递到函数中，则应使用 指针
	// 声明waitGroup
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		i := i

		// 将 worker 调用包装在一个闭包中，可以确保通知 WaitGroup 此工作线程已完成。
		//这样，worker 线程本身就不必知道其执行中涉及的并发原语。
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	wg.Wait()

}
