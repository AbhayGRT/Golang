/*
address and dereference
operators
• & operator - The address of a variable can be obtained
by preceding the name of a variable with an ampersand sign
(&), known as address-of operator.
• * operator - It is known as the dereference operator. When
placed before an address, it returns the value at that address.
*/

package main
import "fmt"

func main() {
	i := 10
	fmt.Printf("%T %v \n", &i, &i)
	fmt.Printf("%T %v \n", *(&i), *(&i))
}