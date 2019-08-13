package main

import "fmt"

func main() {
	//创建容量为5的切片
	x := make([]int, 0, 5)
	for i := 0; i < 8; i++ {
		//追加数据，当超出容量限制时，自动分配更大的存储空间
		x = append(x, i)
	}
	fmt.Println(x)
}
