package main

import "fmt"

/*
	Slice 是 Go 中一个关键的数据类型，是一个比数组更加强大的序列接口
	不像数组，slice 的类型仅由它所包含的元素决定（不像数组中还需要元素的个数）。要创建一个长度非零的空slice，\
	需要使用内建的方法 make。这里我们创建了一个长度为3的 string 类型 slice（初始化为零值）。

	More to read： https://blog.golang.org/go-slices-usage-and-internals
*/

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

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
