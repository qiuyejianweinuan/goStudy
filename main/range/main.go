package main

import "fmt"

func main() {

	nums := []int{2, 3, 4}
	// 声明初始化 sum
	sum := 0
	// 循环累加数组中的数
	for _, num := range nums {
		sum += num
	}
	//打印数组的和 sum: 9
	fmt.Println("sum:", sum)

	// 按照下标输出数组中值为3的数组下标
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// 声明初始化一个map集合
	//a -> apple
	//b -> banana
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		//   %s	the uninterpreted bytes of the string or slice
		fmt.Printf("%s -> %s\n", k, v)
	}

	// 循环输出
	//key:a   value: apple
	//key:b   value: banana
	for k, v := range kvs {
		fmt.Print("key:", k, "\t")
		fmt.Println("value:", v)
	}

	// 循环输出字符串字符的 值 和 对应的ASCII码
	//0 103
	//1 111
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
