package main

import (
	"fmt"
	"math"
)

/* 接口 是方法特征的命名集合。 */

// 定义 geometry 接口
type geometry interface {
	area() float64
	perim() float64
}

// 定义 两种结构体
type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

// 实现 rect的接口方法
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2 * r.width * r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
