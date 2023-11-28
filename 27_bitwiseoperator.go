package main
import "fmt"

// Bitwise Operators

func main() {

	// bitwise AND (&)
	// takes two numbers as operands and does AND on every bit of two numbers.
	var x, y int = 12, 25
	z := x & y
	fmt.Println(z)

	// bitwise OR (|)
	// takes two numbers as operands and does OR on every bit of two numbers.
	var a, b int = 12, 25
	c := a | b
	fmt.Println(c)

	// bitwise XOR (^)
	// takes two numbers as operands and does XOR on every bit of two numbers.
	// The result of XOR is 1 if the two bits are opposite.
	var xa, yb int = 12, 25
	zc := xa ^ yb
	fmt.Println(zc)

	// left shift (<<)
	// shifts all bits towards left by a certain number of specified bits.
	// The bit positions that have been vacated by the left shift operator are filled with 0.
	var i int = 212
	zi := i << 1
	fmt.Println(zi)

	// right shift (>>)
	// shifts all bits towards right by a certain number of specified bits.
	// excess bits shifted off to the right are discarded.
	var j int = 212
	zj := j >> 2
	fmt.Println(zj)
}
