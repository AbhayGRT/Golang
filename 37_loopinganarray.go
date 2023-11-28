package main
import "fmt"

func main()  {
	// var grades [5] int
	// fmt.Println(grades)

	var grade [5] string = [5] string {"Physics","Chemistry","Maths","Biology","Psychology"}
	fmt.Println(grade)

	for i := 0;i < len(grade);i++{
		fmt.Println(grade[i])
	}

	// Different for loop
	for index, element := range grade {
		fmt.Println(index, "=>", element)
		}

	// multidimensional array
	arr := [3][2]int{{2, 4}, {4, 16}, {8, 64}}
	fmt.Println(arr[2][1])
}