package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 原子计数器
// Go中最重要的状态管理机制是依赖通道间的通信来完成的.
// 这节是使用sync/atomic包在多个协程中进行原子计数
func main() {

	// 我们将使用一个无符号整型（永远是正整数）变量来表示这个计数器。
	var ops uint64
	var ops1 uint64

	// WaitGroup 帮助我们等待所有协程完成它们的工作。
	var wg sync.WaitGroup

	// 我们会启动 50 个协程，并且每个协程会将计数器递增 1000 次。
	for i := 0; i < 50; i++ {
		wg.Add(1)
		// 使用 AddUint64 来让计数器自动增加， 使用 & 语法给定 ops 的内存地址。
		go func() {
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&ops, 1)
				ops1++
			}
			wg.Done()
		}()
	}
	// 等待，直到所有协程完成。
	wg.Wait()
	//
	fmt.Println("ops:", ops)
	fmt.Println("not atomic ops1:", ops1)
}
