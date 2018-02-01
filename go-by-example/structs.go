package main
import(
	"fmt"
)
type person struct {
	name string
	age int
}

func main() {
	// 使用这个语法创建了一个新的结构体元素。
	fmt.Println(person{"Bob", 20})

	// 你可以在初始化一个结构体元素时指定字段名字
	fmt.Println(person{name: "Cky", age: 24})

	// 省略的字段将被初始化为零值。
	fmt.Println(person{name: "Cky2"})

	// `&` 前缀生成一个结构体指针。
	fmt.Println(&person{name: "cky3", age:13})

	// 使用点来访问结构体字段
	s := person{name: "cky4", age: 14}
	fmt.Println(s.name)

	// 也可以对结构体指针使用`.` - 指针会被自动解引用。
	sp := &s
	fmt.Println(sp.age)

	// 结构体是可变的。
	sp.age = 51
	fmt.Println(sp.age)
}