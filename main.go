package main

import "fmt"

func main() {
	// 1: /problems/two-sum/
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Printf("Two sum: %v\n", result)
}
