package main
import "fmt"

func main() {
	x := func(l int, b int) int {
		return l * b
	}

	fmt.Println(x(20, 30))

	y := func(l int, b int) int {
		return l * b
	}(20, 30)

	fmt.Println(y)

}