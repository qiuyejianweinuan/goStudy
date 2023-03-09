package main

import (
	"fmt"
	"time"
)

// 定时器是未来某一时刻执行一次的打点器,
func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()
	time.Sleep(1400 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("ticker stopped")
}
