package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Print(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}

/*
	Go 支持字符、字符串、布尔和数值 常量 。
	const 用于声明一个常量。
	const 语句可以出现在任何 var 语句可以出现的地方
	常数表达式可以执行任意精度的运算
	数值型常量是没有确定的类型的，直到它们被给定了一个类型，比如说一次显示的类型转化。
	当上下文需要时，一个数可以被给定一个类型，比如变量赋值或者函数调用。举个例子，这里的 math.Sin函数需要一个 float64 的参数。
*/
