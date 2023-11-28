package main
import "fmt"
/*
the continue statement
skips the current iteration of
loop and continues with the
next iteration.
*/
func main() {

	for i := 1; i <= 5; i++ {
		if i == 3 {
		continue
	}
		fmt.Println(i)
	}

}