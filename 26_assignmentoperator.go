package main
import "fmt"

// Assignment Operators

func main() {

	var x int = 10
	var y int
	y = x
	fmt.Println(y)

	var a, b int = 10, 20
	a += b
	fmt.Println(a)

	var p, q int = 10, 20
	p -= q
	fmt.Println(p)

	var m, n int = 10, 20
	m *= n
	fmt.Println(m)

	var s, t int = 200, 20
	s /= t
	fmt.Println(s)

	var i, j int = 210, 20
	i %= j
	fmt.Println(i)


}
