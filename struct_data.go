package main

import "fmt"

/*
结构体struct可匿名嵌入其他类型
*/

type user struct {
	//结构体类型
	name string
	age  byte
}

type manager struct {
	//匿名嵌入其他类型
	user
	title string
}

func main() {

	var m manager
	m.name = "tom"
	m.age = 29
	m.title = "cto"
	fmt.Println(m)

	var x X
	x.inc()
	println(x)

	//调用 user.ToString()
	println(m.ToString())
}

//可以为当前包内的任意类型定义方法

type X int

//名称前的参数称作receiver，作用类似 python self
func (x *X) inc() {
	*x++
}

//还可以直接调用匿名字段的方法，这种方式可实现与继承类似的功能
func (u manager) ToString() string {
	return fmt.Sprintf("%+v", u)
}
