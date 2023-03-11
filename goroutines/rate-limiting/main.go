package main

import (
	"fmt"
	"time"
)

// 本例子类似于 sentinel 的服务限流
// 速率限制 是控制服务资源利用和质量的重要机制。
//基于协程、通道和打点器，Go 优雅的支持速率限制

func main() {

	// 此例子分为两部门，第一部门模拟为每一个请求都加上200ms限流控制
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	//等待200ms打点器
	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		//等待打点器 limiter 通道每 200ms 接收一个值。 这是我们任务速率限制的调度器。
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// 有时候我们可能希望在速率限制方案中允许短暂的并发请求，并同时保留总体速率限制。
	//我们可以通过缓冲通道来完成此任务。 burstyLimiter 通道允许最多 3 个爆发（bursts）事件。
	burstyLimiter := make(chan time.Time, 3)

	// 先为此通道加上立即时间
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
