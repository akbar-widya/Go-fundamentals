package main

import "fmt"

func minMax(nums []int) (int, int) {
    // Returning dummy data for this example
	return nums[0], nums[1]
}

func main() {
	min, _ := minMax([]int{3, 1, 4, 1, 5})

	fmt.Println("Min:", min)
}