package main
import "fmt"

func main()  {
	// var grades [5] int
	// fmt.Println(grades)

	var grade [5] string = [5] string {"Physics","Chemistry","Maths","Biology","Psychology"}
	fmt.Println(grade)

	grades := [3]int{10, 20, 30}
	fmt.Println(grades)

	grads := [...]int{10, 20, 30}
	fmt.Println(grads)


	// array length
	var fruits [2]string = [2]string{"apples", "oranges"}
	fmt.Println(len(fruits))

	// array indexing
	var fruit [5]string = [5]string{"apples","oranges", "grapes", "mango", "papaya"}
	fmt.Println(fruit[2])
	fruit[1] = "strawberry"
	fmt.Println(fruit)
}