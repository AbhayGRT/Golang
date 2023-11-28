package main

import (
	"fmt"
	"reflect"
)

func main()  {
	fmt.Printf("Type: %v \n", reflect.TypeOf(1000))
	fmt.Printf("Type: %v \n", reflect.TypeOf(455.5))
	fmt.Printf("Type: %v \n", reflect.TypeOf("Abhay"))
	fmt.Printf("Type: %v \n", reflect.TypeOf(false))
}