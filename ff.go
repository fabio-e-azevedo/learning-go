package main
import (
    "fmt"
)

func removeElement(nums []int, val int) int {
    j := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] != val {
    	    if i != j {
    	        nums[j] = nums[i]
		nums[i] = val
		fmt.Println(nums)
    	    }
    	    j++
        }
    }
    nums = nums[:j]
    return len(nums)
}

func main() {
    nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
    fmt.Println(removeElement(nums, 2))
}
