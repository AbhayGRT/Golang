package main
import "fmt"

/*
fallthrough
• The fallthrough
keyword is used in switch-
case to force the
execution flow to fall
through the successive
case block.
*/

func main() {
	var i int = 10
	switch i {
		case -5:
			fmt.Println("-5")
		case 10:
			fmt.Println("10")
			fallthrough
		case 20:
			fmt.Println("20")
			fallthrough
		case 70:
			fmt.Println("20")
			
		default:
		fmt.Println("default")
		}
	}