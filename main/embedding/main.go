package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// 当一个结构体中含有其他结构体，那么此结构体可直接访问被包含结构体的字段
// 一个 container 嵌入 了一个 base. 一个嵌入看起来像一个没有名字的字段
type container struct {
	base
	str string
}

// 内嵌结构体
func main() {

	//当创建含有嵌入的结构体，必须对嵌入进行显式的初始化；
	//在这里使用嵌入的类型当作字段的名字
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	//我们可以直接在 co 上访问 base 中定义的字段, 例如： co.num.
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	//可以直接访问或者写出嵌入结构体(类型)的完整lying
	//或者，我们可以使用嵌入的类型名拼写出完整的路径
	fmt.Println("also num:", co.base.num)

	//由于 container 嵌入了 base，因此base的方法 也成为了 container 的方法。
	//在这里我们直接在 co 上 调用了一个从 base 嵌入的方法
	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("describer:", d.describe())
}
