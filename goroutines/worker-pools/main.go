package main

import (
	"fmt"
	"time"
)

// 这是 worker 程序，我们会并发的运行多个 worker。
// worker 将在 jobs 频道上接收工作，并在 results 上发送相应的结果。
// 每个 worker 我们都会 sleep 一秒钟，以模拟一项昂贵的（耗时一秒钟的）任务。
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second / 2)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	// 启动3个worker完成5个任务
	const numJobs = 5
	// 模拟工作
	jobs := make(chan int, numJobs)
	//
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// 运行程序，显示 5 个任务被多个 worker 执行。
	//尽管所有的工作总共要花费 5 秒钟，但该程序只花了 2 秒钟， 因为 3 个 worker 是并行的。
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
