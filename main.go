package main

import "fmt"

func main() {
	twoSum()
	lengthOfLongestSubstring()
	longestPalindromeSubstring()
	convert()
	reverse()
}

// 1: /problems/two-sum/
func twoSum() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := Leetcode{}.twoSum(nums, target)
	fmt.Println("Two sum:", result)
}

// 3: /problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring() {
	s := "abcdefabc"
	result := Leetcode{}.lengthOfLongestSubstring(s)
	fmt.Println("Length of longest substring:", result)
}

// 5: /problems/longest-palindromic-substring/
func longestPalindromeSubstring() {
	s := "babad"
	result := Leetcode{}.longestPalindromeSubstring(s)
	fmt.Println("Longest palindromic substring:", result)
}

// 6: /problems/zigzag-conversion/
func convert() {
	s := "PAYPALISHIRING"
	numRows := 3
	result := Leetcode{}.convert(s, numRows)
	fmt.Println("Zigzag conversion:", result)
}

// 7: /problems/reverse-integer/
func reverse() {
	x := -2147483648
	result := Leetcode{}.reverse(x)
	fmt.Println("Reverse integer:", result)
}
