package main

import "fmt"

// 1: /problems/two-sum/
func twoSum() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := Leetcode{}.twoSum(nums, target)
	fmt.Println("Two sum:", result)
}

// 2: /problems/add-two-numbers/
func addTwoNumbers() {
	l1 := sliceToLinkedList([]int{9, 9, 9, 9, 9, 9, 9})
	l2 := sliceToLinkedList([]int{9, 9, 9, 9})
	resultLinkedList := Leetcode{}.addTwoNumbers(l1, l2)
	result := linkedListToSlice(resultLinkedList)
	fmt.Println("Adding two linked list numbers:", result)
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

// 8: /problems/string-to-integer-atoi/
func myAtoi() {
	s := " -273 degree "
	result := Leetcode{}.myAtoi(s)
	fmt.Println("String to integer (atoi):", result)
}

// 9: /problems/palindrome-number/
func isPalindrome() {
	x := 121
	result := Leetcode{}.isPalindrome(x)
	fmt.Println("Is palindrome:", result)
}

// 10: /problems/regular-expression-matching/
func isMatch() {
	s := "aa"
	p := "a*"
	result := Leetcode{}.isMatch(s, p)
	fmt.Println("Regular expression matching:", result)
}

// 11: /problems/container-with-most-water/
func maxArea() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	result := Leetcode{}.maxArea(height)
	fmt.Println("Container with most water:", result)
}

// 12: /problems/integer-to-roman/
func intToRoman() {
	num := 2024
	result := Leetcode{}.intToRoman(num)
	fmt.Println("Integer to roman:", result)
}

// 13: /problems/roman-to-integer/
func romanToInt() {
	s := "MMXXIV"
	result := Leetcode{}.romanToInt(s)
	fmt.Println("Roman to integer:", result)
}

// 14: /problems/longest-common-prefix/
func longestCommonPrefix() {
	strs := []string{"flower", "flow", "flight"}
	result := Leetcode{}.longestCommonPrefix(strs)
	fmt.Println("Longest common prefix:", result)
}

// 15: /problems/3sum
func threeSum() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	result := Leetcode{}.threeSum(nums)
	fmt.Println("Three sum:", result)
}

// 16: /problems/3sum-closest/
func threeSumClosest() {
	nums := []int{4, 0, 5, -5, 3, 3, 0, -4, -5}
	target := -2
	result := Leetcode{}.threeSumClosest(nums, target)
	fmt.Println("Three sum closest:", result)
}

// 17: /problems/letter-combinations-of-a-phone-number/
func letterCombination() {
	digits := "38"
	result := Leetcode{}.letterCombinations(digits)
	fmt.Println("Letter Combinations:", result)
}

// 18: /problems/4sum/
func fourSum() {
	nums := []int{0, 0, 0, 1000000000, 1000000000, 1000000000, 1000000000}
	target := 1000000000
	result := Leetcode{}.fourSum(nums, target)
	fmt.Println("Four Sum:", result)
}

// 19: /problems/remove-nth-node-from-end-of-list/
func removeNthFromEnd() {
	head := sliceToLinkedList([]int{1, 2, 3, 4, 5})
	n := 2
	resultLinkedList := Leetcode{}.removeNthFromEnd(head, n)
	result := linkedListToSlice(resultLinkedList)
	fmt.Println("Remove nth node from end of list:", result)
}

// 20: /problems/valid-parentheses/
func isValid() {
	s := "()[]{}"
	result := Leetcode{}.isValid(s)
	fmt.Println("Valid parentheses:", result)
}

// 21: /problems/merge-two-sorted-lists/
func mergeTwoLists() {
	l1 := sliceToLinkedList([]int{1, 2, 4})
	l2 := sliceToLinkedList([]int{1, 3, 4})
	resultLinkedList := Leetcode{}.mergeTwoLists(l1, l2)
	result := linkedListToSlice(resultLinkedList)
	fmt.Println("Merge two lists:", result)
}

// 22: /problems/generate-parentheses/
func generateParenthesis() {
	n := 3
	result := Leetcode{}.generateParenthesis(n)
	fmt.Println("Generate parentheses:", result)
}

func main() {
	twoSum()
	addTwoNumbers()
	lengthOfLongestSubstring()
	longestPalindromeSubstring()
	convert()
	reverse()
	myAtoi()
	isPalindrome()
	isMatch()
	maxArea()
	intToRoman()
	romanToInt()
	longestCommonPrefix()
	threeSum()
	threeSumClosest()
	letterCombination()
	fourSum()
	removeNthFromEnd()
	isValid()
	mergeTwoLists()
	generateParenthesis()
}
