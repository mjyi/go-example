package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 20})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Ann", age: 30})

	s := person{name: "Sean", age: 49}
	fmt.Println(s.name)

	ss := s
	ss.age = 54
	fmt.Println("ss", ss)
	fmt.Println("s", s)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 53
	fmt.Println(s.age)
}
