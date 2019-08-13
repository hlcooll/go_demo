package main

/*
循环语句
*/
import "fmt"

func ifelse(aa int, bb int) {
	if aa < bb {
		fmt.Printf("a 小于 bb\n")
	} else {
		fmt.Printf("a 大于 bb\n")
	}
}

func ifif(aa int, bb int) {
	if aa == 100 {
		if bb == 200 {
			fmt.Printf("aa 的值为 100 ，bb的值为200\n")
		}
	} else {
		fmt.Printf("aa 不为100\n")
	}

	fmt.Printf("aa 值为：%d\n", aa)
	fmt.Printf("bb 值为：%d\n", bb)
}

func ifswitch(grade string, marks int) {
	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "G"
	}
	switch {
	case grade == "A":
		fmt.Printf("优秀!\n")
	case grade == "B", grade == "C":
		fmt.Printf("良好\n")
	case grade == "D":
		fmt.Printf("及格\n")
	case grade == "F":
		fmt.Printf("不及格\n")
	default:
		fmt.Printf("差\n")
	}
	fmt.Printf("你的等级是 %s\n", grade)
}

func ifselect() {
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := (<-c3): // same as: i3, ok := <-c3
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}
}
func jiajian(a int) {

	a -= 10 * 2

	fmt.Printf("a= %d\n", a)

}

func main() {
	//ifelse(12,50)
	//	//ifelse(70,50)
	//	//ifif(990,200)
	//	//ifswitch("M",3)
	//	//ifselect()
	jiajian(2)
}
