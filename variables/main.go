package main

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

type Generics1 interface {
	Struct1 | Struct2 | Struct3 | Struct4
}

type Struct1 struct {
	value int32
}

type Struct2 struct {
	value1 bool  // 1 byte *nesse caso o type bool ocupa 4 byte (3 byte padding)
	value2 int32 // 4 bytes
	value3 int32 // 4 bytes
} // valor total de 12 bytes

type Struct3 struct {
	value1 bool  // 1 byte *nesse caso o type bool ocupa 4 byte (3 byte padding)
	value2 int32 // 4 bytes
	value3 int32 // 4 bytes
	value4 int64 // 8 bytes
} // valor total de 24 bytes pq tem q ser multiplo do maior valor

type Struct4 struct {
	value1 int32 // 4 bytes (4 bytes padding)
	value2 int64 // 8 bytes
	value3 bool  // 1 byte *nesse caso o type bool ocupa 8 byte (7 bytes padding)
} // valor total de 24 bytes

func main() {
	var variable1 bool
	variable2 := true

	var variable3 int
	var variable4 int8
	var variable5 int16
	var variable6 int32
	var variable7 int64
	var variable8 string

	var array1 [5]int64
	var slice1 []int64

	totalMap := make(map[string]interface{})

	totalMap["variable1"] = &variable1
	totalMap["variable2"] = &variable2
	totalMap["variable3"] = &variable3
	totalMap["variable4"] = &variable4
	totalMap["variable5"] = &variable5
	totalMap["variable6"] = &variable6
	totalMap["variable7"] = &variable7
	totalMap["variable8"] = &variable8

	totalMap["Struct1"] = Struct1{}
	totalMap["Struct2"] = Struct2{}
	totalMap["Struct3"] = Struct3{}
	totalMap["Struct4"] = Struct4{}

	for chave, valor := range totalMap {
		switch v := valor.(type) {
		case *int:
			fmt.Printf("%s - type int           = %v bytes\n", chave, unsafe.Sizeof(*v))
		case *int8:
			fmt.Printf("%s - type int8          = %v byte\n", chave, unsafe.Sizeof(*v))
		case *int16:
			fmt.Printf("%s - type int16         = %v bytes\n", chave, unsafe.Sizeof(*v))
		case *int32:
			fmt.Printf("%s - type int32         = %v bytes\n", chave, unsafe.Sizeof(*v))
		case *int64:
			fmt.Printf("%s - type int64         = %v bytes\n", chave, unsafe.Sizeof(*v))
		case *string:
			fmt.Printf("%s - type string        = %v bytes\n", chave, unsafe.Sizeof(*v))
		case *bool:
			fmt.Printf("%s - type bool          = %v byte\n", chave, unsafe.Sizeof(*v))
		case interface{}:
			switch {
			case strings.Contains(chave, "Struct"):
				fmt.Printf("%s   - type Struct        = %v bytes\n", chave, reflect.TypeOf(v).Size())
			}
		default:
			fmt.Printf("Tipo de dado não reconhecido para %s %s\n", chave, v)
		}
	}

	fmt.Printf("array1    - type [3]int64      = %v bytes\n", unsafe.Sizeof(array1))
	fmt.Printf("slice1    - type []int64       = %v bytes\n", unsafe.Sizeof(slice1))
}
