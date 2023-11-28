package main
import (
"fmt"
"strconv")

/*
strconv package
Itoa()
• converts integer to string
• returns one value – string formed with the given integer.
*/

func main() {
	var i int = 42
	var s string = strconv.Itoa(i) // convert int to string
	fmt.Printf("%q", s)
}