package main
import "fmt"

// functions without return type
func printGreeeting(str string) {
	fmt.Println("Hey there,", str)
}
// functions with return type
func addNumbers(a int, b int) int {
	sum := a + b
	return sum
	}

// returning multiple values
func operation(a int, b int) (int, int){
	sum := a + b
	diff := a - b
	return sum, diff
}

func main() {
	printGreeeting("Abhay")

	sumOfNumbers := addNumbers(2, 3)
	fmt.Println(sumOfNumbers)

	sum, difference := operation(20, 10)
	fmt.Println(sum, difference)
}