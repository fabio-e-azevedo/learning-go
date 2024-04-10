package main

import "fmt"

func main() {
	//nums := []int{2, 2, 3, 2}
	nums := []int{0, 1, 0, 1, 0, 1, 99}

	fmt.Println(singleNumber(nums))
}

func singleNumber(nums []int) int {
	var number int
	for _, n := range nums {
		number = 0
		for i := range nums {
			if n == nums[i] {
				number++
			}
		}
		if number == 1 {
			return n
		}
	}
	return 0
}

// func singleNumber(nums []int) int {
// 	var ones int = 0
// 	var twos int = 0

// 	for i := 0; i < len(nums); i++ {
// 		var number int = nums[i]
// 		ones ^= (number & ^twos)
// 		twos ^= (number & ^ones)
// 	}

// 	return ones
// }
