package main

import "fmt"

func main() {
	twoSum()
	lengthOfLongestSubstring()
}

func twoSum() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := Leetcode{}.twoSum(nums, target)
	fmt.Println("Two sum:", result)
}

func lengthOfLongestSubstring() {
	s := "abcdefabc"
	result := Leetcode{}.lengthOfLongestSubstring(s)
	fmt.Println("Length of longest substring:", result)
}
