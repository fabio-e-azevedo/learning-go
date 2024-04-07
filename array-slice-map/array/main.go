package main

import (
	"fmt"
	"reflect"
)

func main() {
	//Declaration 1
	//var arr = [5]int{1, 2, 3, 4, 5}
	var arr [5]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5
	//arr := [5]int{1,2,3,4,5}

	fmt.Println(arr)
	fmt.Printf("Variable arr is Type %s: %T\n", reflect.TypeOf(arr).Kind(), arr)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)
	fmt.Printf("Variable names is Type %s: %T\n", reflect.TypeOf(names).Kind(), names)
}
