package main
import "fmt"

// Logical Operators
// • used to determine the logic between variables or values.
// • common logical comparisons -
// • Are two variables both true ?
// • Does either of two expressions evaluate to true?

func main() {

// && (AND)
// • returns true if both the statements are true.
// • returns false when either of the statement is false.
	var x int = 10
	fmt.Println((x < 100) && (x < 200))
	fmt.Println((x < 300) && (x < 0))

// || (OR)
// • returns true if one of the statement is true.
// • returns false when both statements are false.
	var y int = 10
	fmt.Println((y < 100) || (y < 200))
	fmt.Println((y < 300) || (y < 0))

// !
// • unary operator.
// • Reverses the result, returns false if the expression evaluates to true and vice versa.
	var a, b int = 10, 20
	fmt.Println(!(a > b))
	fmt.Println(!(true))
	fmt.Println(!(false))

}
