package main
import(
	"fmt"
)

func values () (int, int) {
	return 3, 7
}

func main() {
	a, b := values()
	fmt.Println(a, b)

	_, c := values()
	fmt.Println(c)
}