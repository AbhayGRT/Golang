package main
import "fmt"
func main() {

	arr := [4]int{10, 20, 30, 40}
	slice := arr[1:3]

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice = append(slice, 900, -90, 50)

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))


	src_slice := []int{10, 20, 30, 40, 50}
	dest_slice := make([]int, 3)
	num := copy(dest_slice, src_slice)
	fmt.Println(dest_slice)
	fmt.Println("Number of elements copied: ", num)
}