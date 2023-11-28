package main
import "fmt"

func main() {
	var grades int = 42
	var message string = "hello world"
	var isCheck bool = true
	var amount float64 = 455.6

	fmt.Printf("variable grades = %v is of type %T \n", grades, grades)
	fmt.Printf("variable message = '%v' is of type %T \n", message, message)
	fmt.Printf("variable isCheck = '%v' is of type %T \n", isCheck, isCheck)
	fmt.Printf("variable amount = %v is of type %T \n", amount, amount)
}