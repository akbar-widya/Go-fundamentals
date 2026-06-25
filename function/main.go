package main

import "fmt"

func minMax(nums []int) (min int, max int) {
    min = nums[0]
	max = nums[0]

	for _, n := range nums {	// n is the value from index 0, then the program is loop with the next index's value
		if n < min { min = n }
		if n > max { max = n }
	}
	return
}

func main() {
	min, max := minMax([]int{3, 1, 4, 1, 5})

	fmt.Println("Min:", min, "Max:", max)
}