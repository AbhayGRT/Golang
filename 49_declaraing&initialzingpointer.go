package main
import "fmt"

func main() {

	// declaring a pointer
	var i *int
	var r *string
	fmt.Println(i)
	fmt.Println(r)

	// initialzing a pointer
	s := "hello"
	var b *string = &s
	fmt.Println(b)
	var a = &s
	fmt.Println(a)
	c := &s
	fmt.Println(c)
}