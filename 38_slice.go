package main
import "fmt"
func main() {
	slice := []int{10, 20, 30}
	fmt.Println(slice)

	arr := [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	slice_1 := arr[1:8]
	fmt.Println(slice_1)

	sub_slice := slice[0:3]
	fmt.Println(sub_slice)


	// Empty slice
	sliced := make([]int, 5, 8)
	fmt.Println(sliced)
	fmt.Println(len(sliced))
	fmt.Println(cap(sliced))


	arr1 := [5]int{10, 20, 30, 40, 50}
	slices := arr1[:3]
	fmt.Println(arr1)
	fmt.Println(slices)
	
	slices[1] = 9000
	fmt.Println("after modification")
	fmt.Println(arr1)
	fmt.Println(slices)
}