package main

import "fmt"

//Go 协程 在执行上来说是轻量级的线程。
func f(from string) {
	for i := 0; i < 13; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	go f("goroutine")

	// 你为匿名函数启动一个 Go 协程。
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}

/*
	将首先看到阻塞式调用的输出，然后是两个 Go 协程的交替输出。
	这种交替的情况表示 Go 运行时是以异步的方式运行协程的。
*/
