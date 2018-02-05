// Go 的 `sort` 包实现了内置和用户自定义数据类型的排序功能。
// 我们首先关注内置数据类型的排序。
package main
import(
	"sort"
	"fmt"
)

func main() {
	// 排序方法是正对内置数据类型的；这里是一个字符串的例子。
	// 注意排序是原地更新的，所以他会改变给定的序列并且不返回一个新值。
	strs := []string{"c", "y", "k"}
	sort.Strings(strs)
	fmt.Println("strings:", strs)

	ints := []int{4,2,1}
	sort.Ints(ints)
	fmt.Println("ints:", ints)

	// 我们也可以使用 `sort` 来检查一个序列是不是已经是排好序的。
	s := sort.IntsAreSorted(ints)
	fmt.Println("sorted:", s)
}