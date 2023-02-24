package main

import "fmt"

func multipleReturn() (int, int) {
	return 3, 4
}

func main() {

	a, b := multipleReturn()
	fmt.Println(a)
	fmt.Println(b)

	_, c := multipleReturn()
	fmt.Println(c)
}
