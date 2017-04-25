package main

import "fmt"

/*
	map 是 Go 内置关联数据类型（在一些其他的语言中称为哈希 或者字典 ）。
	make:make(map[key-type]val-type). 创建map
	使用典型的 make[key] = val 语法来设置键值对。
	内建的 delete 可以从一个 map 中移除键值对
	当从一个 map 中取值时，可选的第二返回值指示这个键是在这个 map 中。这可以用来消除键不存在和键有零值，像 0 或者 "" 而产生的歧义。
*/

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)
	fmt.Println("len", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
