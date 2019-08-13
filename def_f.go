package main

/*
函数
*/
import (
	"errors"
)

//1
func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}

	return a / b, nil
}

//2
func test(x int) func() {
	//常用来释放资源、解除锁定，或执行一些清理操作
	//可定义多个defer，按FILO顺序执行
	defer println("defer定义延迟调用，无论函数是否出错，它都确保结束前被调用")
	return func() {
		println(x)
	}
}

func main() {
	//1
	//a,b :=10,0
	//c,err:=div(a,b)
	//fmt.Println(c,err)
	//2
	x := 100
	f := test(x)
	f()

}
