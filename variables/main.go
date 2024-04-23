package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var variable1 bool
	variable2 := true

	var variable3 int
	var variable4 int8
	var variable5 int16
	var variable6 int32
	var variable7 int64

	var array1 [5]int64
	var slice1 []int64

	var struct1 struct {
		value int32
	}

	var struct2 struct {
		value1 bool  // 1 byte *nesse caso o type bool ocupa 4 byte (3 byte padding)
		value2 int32 // 4 bytes
		value3 int32 // 4 bytes
	} // valor total de 12 bytes

	var struct3 struct {
		value1 bool  // 1 byte *nesse caso o type bool ocupa 4 byte (3 byte padding)
		value2 int32 // 4 bytes
		value3 int32 // 4 bytes
		value4 int64 // 8 bytes
	} // valor total de 24 bytes pq tem q ser multiplo do maior valor

	var struct4 struct {
		value1 int32 // 4 bytes (4 bytes padding)
		value2 int64 // 8 bytes
		value3 bool  // 1 byte *nesse caso o type bool ocupa 8 byte (7 bytes padding)
	} // valor total de 24 bytes

	fmt.Printf("variable1 - type bool          = %v byte\n", unsafe.Sizeof(variable1))
	fmt.Printf("variable2 - type bool          = %v byte\n", unsafe.Sizeof(variable2))
	fmt.Printf("variable3 - type int           = %v bytes\n", unsafe.Sizeof(variable3))
	fmt.Printf("variable4 - type int8          = %v byte\n", unsafe.Sizeof(variable4))
	fmt.Printf("variable5 - type int16         = %v bytes\n", unsafe.Sizeof(variable5))
	fmt.Printf("variable6 - type int32         = %v bytes\n", unsafe.Sizeof(variable6))
	fmt.Printf("variable7 - type int64         = %v bytes\n", unsafe.Sizeof(variable7))
	fmt.Printf("array1    - type [3]int64      = %v bytes\n", unsafe.Sizeof(array1))
	fmt.Printf("slice1    - type []int64       = %v bytes\n", unsafe.Sizeof(slice1))
	fmt.Printf("struct1   - type struct{}      = %v bytes\n", unsafe.Sizeof(struct1))
	fmt.Printf("struct2   - type struct{}      = %v bytes\n", unsafe.Sizeof(struct2))
	fmt.Printf("struct3   - type struct{}      = %v bytes\n", unsafe.Sizeof(struct3))
	fmt.Printf("struct4   - type struct{}      = %v bytes\n", unsafe.Sizeof(struct4))
}
