package main

import "fmt"

func main() {
	f := func(x, y int) int {
		return x + y
	}
	fmt.Println(f(1, 2))
}
