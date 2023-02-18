package main

import "fmt"

func main() {

	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	//切片操作
	l = s[:5]
	fmt.Println("sl2:", l)

	//切片
	l = s[2:]
	fmt.Println("sl3:", l)

	//初始化slice
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	//初始化长度为3的数组的切片
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		//每次切片内数组长度加 1
		innerLen := i + 1
		//定义切片内数组
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	twoA := make([][]int, 2)
	for i := 0; i < 2; i++ {
		innerLen := i + 1
		twoA[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoA[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
	fmt.Println("2A: ", twoA)

}
