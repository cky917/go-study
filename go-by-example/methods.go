package main
import(
	"fmt"
)
type rect struct {
	width, height int
}

func (r * rect) area() int {
	return r.width * r.height
}

// 可以为值类型或者指针类型的接收器定义方法。这里是一个
// 值类型接收器的例子。
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 3}

	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	// Go 自动处理方法调用时的值和指针之间的转化。
	// 你可以使用指针来调用方法来避免在方法调用时产生一个拷贝，或者让方法能够改变接受的数据。
	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())
}