package main
import 
("fmt"
"strconv")
// Atoi()
// • converts string to integer.
// • returns two values – the corresponding integer, error (if any).
func main()  {
	var s string = "200"
	i,err := strconv.Atoi(s)
	fmt.Println(i)
	fmt.Print(err)
}