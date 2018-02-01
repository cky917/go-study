package main
import(
	"fmt"
	"errors"
)
// Go 语言使用一个独立的·明确的返回值来传递错误信息。
// 这与使用异常的 Java 和 Ruby 以及在 C 语言中经常见到的超重的单返回值/错误值相比，
// Go 语言的处理方式能清楚的知道哪个函数返回了错误，并能像调用那些没有出错的函数一样调用。

// 按照惯例，错误通常是最后一个返回值并且是 `error` 类型，一个内建的接口。
func f1(arg int) (int, error) {
	if arg == 42 {
		// `errors.New` 构造一个使用给定的错误信息的基本`error` 值。
		return -1, errors.New("Can't work with 42")
	}
	return arg + 3, nil
}
// 通过实现 `Error` 方法来自定义 `error` 类型是可以得。
// 这里使用自定义错误类型来表示上面的参数错误。

type argError struct {
	arg int
	prob string
}

func (e * argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "Can't work with 42"}
	}
	return arg + 3, nil
}

func main() {
	for _, i := range []int{7, 42} {
		if result, err := f1(i); err != nil {
			fmt.Println("f1 failed:", err)
		} else {
			fmt.Println("f1 worked:", result)
		}
	}

	for _, i := range []int{7, 42} {
		if result, err := f2(i); err != nil {
			fmt.Println("f2 failed:", err)
		} else {
			fmt.Println("f2 worked:", result)
		}
	}

	// 你如果想在程序中使用一个自定义错误类型中的数据，你需要通过类型断言来得到这个错误类型的实例。
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}