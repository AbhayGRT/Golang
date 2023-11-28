package main
import "fmt"

// type Student struct {
// 	name string
// 	rollNo int
// 	marks []int
// 	grades map[string]int
// }

// func main() {
// 	var s Student
// 	fmt.Printf("%+v", s)
// }


type Student struct {
	name string
	rollNo int
	}
	
func main() {

	st := Student{
		name: "Joe",
		rollNo: 12,
	}

	fmt.Printf("%+v", st)
}