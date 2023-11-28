package main
import "fmt"

// Arithmetic Operators
// • used to perform common arithmetic operations, such as addition,
// subtraction, multiplication etc.
// • common comparisons -
// • Does the sum of two numbers equal a particular value?
// • Is the difference between two numbers lesser than a particular value?

func main() {

// + addition
	var a,b string = "foo", "bar"
	fmt.Println(a + b)

// - subtraction
	var c, d float64 = 79.02, 75.66
	fmt.Printf("%.2f", c - d)

// * multiplication
	var e, f int = 12, 2
	fmt.Println(e * f)

// / division
	var x, y int = 24, 2
	fmt.Println(x / y)

// ++ increment (unary operator)
	var i int = 1
	i++
	fmt.Println(i)

// -- decrement (unary operator)
	var u int = 1
	u--
	fmt.Println(u)
}
