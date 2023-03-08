package main

import "fmt"

func main() {
	var n string
	_, err := fmt.Scanln(&n)
	if err != nil {
		return
	}
	fmt.Println(n)
	var m string
	_, err = fmt.Scanf("%v", &m)
	if err != nil {
		return
	}
	fmt.Println(m)
}
