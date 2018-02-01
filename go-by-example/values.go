package main
import (
	"fmt"
)
func main() {
	// 字符串k可以通过 + 连接
	fmt.Println("go" + "lang")
	
	// 整数 浮点数
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}