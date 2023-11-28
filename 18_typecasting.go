package main
import "fmt"
func main() {
	var i int = 90
	var f float64 = float64(i)
	fmt.Printf("%.2f\n", f)


	var x float64 = 45.89
	var y int = int(x)
	fmt.Printf("%v\n", y)
}