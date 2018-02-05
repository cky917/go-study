package main
import(
	"os"
)

func main() {
	// 我们将在真个网站中使用 panic 来检查预期外的错误。
	// 这个是唯一一个为 panic 准备的例子。
	panic("a problem")
	// panic 的一个基本用法就是在一个函数返回了错误值但是我们并不知道（或者不想）处理时终止运行。// 这里是一个在创建一个新文件时返回异常错误时的`panic` 用法。
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}