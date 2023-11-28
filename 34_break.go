package main
import "fmt"

/*
the break statement ends
the loop immediately when
it is encountered.
*/

func main() {
	for i := 1; i <= 5; i++ {
		if i == 3 {
		break
	}
		fmt.Println(i)
	}
}