package main
import (
	"fmt"
)
func main() {
	var a string = "initial"
	fmt.Println(a)

	var b,c int = 1,2
	fmt.Println(b,c)

	// Go 将自动推断已经初始化的变量类型。
	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f:="short"
	fmt.Println(f)
}