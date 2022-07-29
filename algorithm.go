#99乘法表
func main() {
	var i, j int
	for i = 0; i <= 9; i++ {
		for j = 0; j <= 9; j++ {
			fmt.Printf("%d*%d=%d\t", i, j, i*j)
		}
		fmt.Println()
	}
}
