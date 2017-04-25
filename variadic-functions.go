package main

import "fmt"

/*
	可变参数函数。可以用任意数量的参数调用。例如，fmt.Println 是一个常见的变参函数。
	这个函数使用任意数目的 int 作为参数。
*/
func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
