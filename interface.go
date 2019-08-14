package main

import "fmt"

type user1 struct {
	//结构体类型
	name string
	age  byte
}

func (u user1) Print() {
	fmt.Printf("%+v\n", u)
}

type Printer interface {
	Print()
}

func main() {
	var u user1
	u.name = "tttt"
	u.age = 24

	var p Printer = u
	p.Print()
}
