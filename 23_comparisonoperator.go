package main
import "fmt"

// Comparison Operator
// • compare two operands and yield a Boolean value.
// • allow values of the same data type for comparisons
// • common comparisons -
// 	Does one string match another ?
// 	Are two numbers the same ?
// 	Is one number greater than another ?

func main() {
// == equal
	var city string = "Kolkata"
	var city_2 string = "Calcutta"
	fmt.Println(city == city_2)
// != not equal
	var city_3 string = "Kolkata"
	var city_4 string = "Calcutta"
	fmt.Println(city_3 != city_4)
// < less than
	var a, b int = 5, 10
	fmt.Println(a < b)
// <= less than or equal to
	var c, d int = 10, 10
	fmt.Println(c <= d)
// > greater than
	var e, f int = 20, 10
	fmt.Println(e > f)
// >= greater than or equal to
	var g, h int = 20, 20
	fmt.Println(g >= h)
}




