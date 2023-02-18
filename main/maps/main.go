package main

import "fmt"

func main() {

	//golang声明一个map集合
	m := make(map[string]int)

	//为map集合添加键值对
	m["k1"] = 7
	m["k2"] = 13

	//输出  map: map[k1:7 k2:13]
	fmt.Println("map:", m)

	//为v1赋值map集合的k1对应的值
	v1 := m["k1"]

	//输出  v1:  7
	fmt.Println("v1: ", v1)

	// 输出map集合的长度
	fmt.Println("len:", len(m))

	// 删除map集合的k2键
	delete(m, "k2")
	fmt.Println("map:", m)

	// 获取 m 中的 k2是否存在
	_, prs := m["k2"]
	//prs: false
	fmt.Println("prs:", prs)

	// 初始化一个map集合
	n := map[string]int{"foo": 1, "bar": 2}
	//输出  map: map[bar:2 foo:1]
	fmt.Println("map:", n)
}
