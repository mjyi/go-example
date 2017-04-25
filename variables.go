package main

import (
	"fmt"
)

func main() {
	var a string = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "short"
	fmt.Println(f)

}

/*
	在 Go 中，变量被显式声明，并被编译器所用来检查函数调用时的类型正确性
	:= 语句是申明并初始化变量的简写，例如这个例子中的 var f string = "short"。
*/
