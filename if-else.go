package main

import (
	"fmt"
)

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negatvie")
	} else if num < 10 {
		fmt.Println("has 1 digit")
	} else {
		fmt.Print(num, "has mutiple digit")
	}
}

/*
	在条件语句之前可以有一个语句；任何在这里声明的变量都可以在所有的条件分支中使用。
	Go 里没有三目运算符，所以即使你只需要基本的条件判断，你仍需要使用完整的 if 语句
*/
