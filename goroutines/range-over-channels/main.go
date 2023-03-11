package main

import "fmt"

// for 和 range 为基本的数据结构提供迭代功能,我们也可以使用range来遍历通道
func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// 此循环,当queue去所有值取出来的时候遍历结束
	for item := range queue {
		fmt.Println(item)
	}
}
