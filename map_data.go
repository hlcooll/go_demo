package main

/*
将字典map类型内置，可直接从运行时层面获得性能优化
*/
import "fmt"

func main() {
	//创建字典类型对象
	m := make(map[string]int)
	//添加或设置
	m["a"] = 1
	//使用ok-idiom获取值，可知道 key/value是否存在
	//所谓ok-idiom模式，是指在多返回值中用一个名为ok的布尔值来标示操作是否成功，因为很多操作默认返回零值
	x, ok := m["b"]

	fmt.Println(x, ok)
	//删除
	delete(m, "a")
}
