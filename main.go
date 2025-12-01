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

// 23: /problems/merge-k-sorted-lists/
func mergeKLists() {
	lists := []*ListNode{
		sliceToLinkedList([]int{1, 4, 5}),
		sliceToLinkedList([]int{1, 3, 4}),
		sliceToLinkedList([]int{2, 6}),
	}
	resultLinkedList := Leetcode{}.mergeKLists(lists)
	result := linkedListToSlice(resultLinkedList)
	fmt.Println("Merge k sorted lists:", result)
}

// 24: /problems/swap-nodes-in-pairs/
func swapPairs() {
	head := sliceToLinkedList([]int{1, 2, 3, 4, 5, 6})
	resultLinkedList := Leetcode{}.swapPairs(head)
	result := linkedListToSlice(resultLinkedList)
	fmt.Println("Swap nodes in pairs:", result)
}

// 25: /problems/reverse-nodes-in-k-group/
func reverseKGroup() {
	head := sliceToLinkedList([]int{1, 2, 3, 4, 5})
	k := 2
	resultLinkedList := Leetcode{}.reverseKGroup(head, k)
	result := linkedListToSlice(resultLinkedList)
	fmt.Println("Reverse nodes in k group:", result)
}

// 26: /problems/remove-duplicates-from-sorted-array/
func removeDuplicates() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	result := Leetcode{}.removeDuplicates(nums)
	fmt.Println("Remove duplicates from sorted array:", result)
}

// 27: /problems/remove-element/
func removeElement() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	result := Leetcode{}.removeElement(nums, val)
	fmt.Println("Remove element:", result)
}

// 28: /problems/find-the-index-of-the-first-occurrence-in-a-string/
func strStr() {
	haystack := "sadness"
	needle := "sad"
	result := Leetcode{}.strStr(haystack, needle)
	fmt.Println("Find the index of the first occurrence in a string:", result)
}

// 29: /problems/divide-two-integers/
func divide() {
	dividend := 7
	divisor := -3
	result := Leetcode{}.divide(dividend, divisor)
	fmt.Println("Divide two integers:", result)
}

// 30: /problems/substring-with-concatenation-of-all-words/
func findSubstring() {
	s := "barfoothefoobarman"
	words := []string{"foo", "bar"}
	result := Leetcode{}.findSubstring(s, words)
	fmt.Println("Substring with concatenation of all words:", result)
}

// 31: /problems/next-permutation/
func nextPermutation() {
	nums := []int{4, 5, 3, 2, 1}
	Leetcode{}.nextPermutation(nums)
	fmt.Println("Next permutation:", nums)
}

// 32: /problems/longest-valid-parentheses/
func longestValidParentheses() {
	s := ")()())"
	result := Leetcode{}.longestValidParentheses(s)
	fmt.Println("Longest valid parentheses:", result)
}

// 33: /problems/search-in-rotated-sorted-array/
func search() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	result := Leetcode{}.search(nums, target)
	fmt.Println("Search in rotated sorted array:", result)
}

// 34: /problems/find-first-and-last-position-of-element-in-sorted-array/
func searchRange() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	result := Leetcode{}.searchRange(nums, target)
	fmt.Println("Find first and last position of element in sorted array:", result)
}

// 35: /problems/search-insert-position/
func searchInsert() {
	nums := []int{1, 3, 5, 6}
	target := 0
	result := Leetcode{}.searchInsert(nums, target)
	fmt.Println("Search insert position:", result)
}

// 36: /problems/valid-sudoku/
func isValidSudoku() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	result := Leetcode{}.isValidSudoku(board)
	fmt.Println("Valid sudoku:", result)
}

// 37: /problems/sudoku-solver/
func solveSudoku() {
	board := [][]byte{
		{'4', '5', '.', '.', '.', '9', '3', '.', '1'},
		{'.', '.', '.', '.', '8', '.', '9', '.', '.'},
		{'.', '2', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '2', '.', '.', '4', '.', '.', '.'},
		{'.', '.', '8', '.', '.', '.', '.', '.', '3'},
		{'3', '4', '.', '.', '7', '.', '.', '5', '.'},
		{'1', '9', '.', '.', '.', '8', '.', '.', '5'},
		{'.', '.', '3', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '6', '.', '.', '.', '7', '.'},
	}
	Leetcode{}.solveSudoku(board)
	fmt.Println("Sudoku solver: [")
	for _, row := range board {
		fmt.Printf("\t[")
		for i, cell := range row {
			fmt.Print(string(cell))
			if i < len(row)-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println("]")
	}
	fmt.Println("]")
}

// 38: /problems/count-and-say/
func countAndSay() {
	n := 5
	result := Leetcode{}.countAndSay(n)
	fmt.Println("Count and say:", result)
}

// 39: /problems/combination-sum/
func combinationSum() {
	candidates := []int{2, 3, 5}
	target := 8
	result := Leetcode{}.combinationSum(candidates, target)
	fmt.Println("Combination sum:", result)
}

// 40: /problems/combination-sum-ii/
func combinationSum2() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	result := Leetcode{}.combinationSum2(candidates, target)
	fmt.Println("Combination sum 2:", result)
}

// 41: /problems/first-missing-positive/
func firstMissingPositive() {
	nums := []int{3, 4, -1, 1}
	result := Leetcode{}.firstMissingPositive(nums)
	fmt.Println("First missing positive:", result)
}

// 42: /problems/trapping-rain-water/
func trap() {
	height := []int{4, 2, 0, 3, 2, 5}
	result := Leetcode{}.trap(height)
	fmt.Println("Trapping rain water:", result)
}

// 43: /problems/multiply-strings/
func multiply() {
	num1 := "123"
	num2 := "456"
	result := Leetcode{}.multiply(num1, num2)
	fmt.Println("Multiply strings:", result)
}

// 44: /problems/wildcard-matching/
func isMatchWildcard() {
	s := "aa"
	p := "a*"
	result := Leetcode{}.isMatchWildcard(s, p)
	fmt.Println("Wildcard matching:", result)
}

// 45: /problems/jump-game-ii/
func jump() {
	nums := []int{2, 3, 1, 1, 4}
	result := Leetcode{}.jump(nums)
	fmt.Println("Jump game II:", result)
}

// 46: /problems/permutations/
func permute() {
	nums := []int{1, 2, 3}
	result := Leetcode{}.permute(nums)
	fmt.Println("Permutations:", result)
}

// 47: /problems/permutations-ii/
func permuteUnique() {
	nums := []int{1, 1, 2}
	result := Leetcode{}.permuteUnique(nums)
	fmt.Println("Permutations II:", result)
}

// 48: /problems/rotate-image/
func rotate() {
	matrix := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	Leetcode{}.rotate(matrix)
	fmt.Println("Rotate image:", matrix)
}

// 49: /problems/group-anagrams/
func groupAnagrams() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	result := Leetcode{}.groupAnagrams(strs)
	fmt.Println("Group anagrams:", result)
}

// 50: /problems/powx-n/
func myPow() {
	x := 2.0
	n := 10
	result := Leetcode{}.myPow(x, n)
	fmt.Printf("Pow(%v, %v): %v\n", x, n, result)
}

// 51: /problems/n-queens/
func solveNQueens() {
	n := 4
	result := Leetcode{}.solveNQueens(n)
	fmt.Printf("Solve %v-Queens: %v\n", n, result)
}

// 52: /problems/n-queens-ii/
func totalNQueens() {
	n := 4
	result := Leetcode{}.totalNQueens(n)
	fmt.Printf("Total %v-Queens: %v\n", n, result)
}

// 53: /problems/maximum-subarray/
func maxSubArray() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	result := Leetcode{}.maxSubArray(nums)
	fmt.Println("Maximum subarray:", result)
}

// 54: /problems/spiral-matrix/
func spiralOrder() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	result := Leetcode{}.spiralOrder(matrix)
	fmt.Println("Spiral matrix:", result)
}

// 55: /problems/jump-game/
func canJump() {
	nums := []int{3, 2, 1, 0, 4}
	result := Leetcode{}.canJump(nums)
	fmt.Println("Jump game:", result)
}

// 56: /problems/merge-intervals/
func mergeIntervals() {
	intervals := [][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}
	result := Leetcode{}.mergeIntervals(intervals)
	fmt.Println("Merge intervals:", result)
}

// 57: /problems/insert-interval/
func insertInterval() {
	intervals := [][]int{
		{1, 2},
		{3, 5},
		{6, 7},
		{8, 10},
		{12, 16},
	}
	newInterval := []int{4, 8}
	result := Leetcode{}.insertInterval(intervals, newInterval)
	fmt.Println("Insert interval:", result)
}

// 58: /problems/length-of-last-word/
func lengthOfLastWord() {
	s := "Hello World"
	result := Leetcode{}.lengthOfLastWord(s)
	fmt.Println("Length of last word:", result)
}

// 59: /problems/spiral-matrix-ii/
func generateMatrix() {
	n := 3
	result := Leetcode{}.generateMatrix(n)
	fmt.Println("Spiral matrix II:", result)
}

// 60: /problems/permutation-sequence/
func getPermutation() {
	n := 4
	k := 9
	result := Leetcode{}.getPermutation(n, k)
	fmt.Println("Permutation sequence:", result)
}

// 61: /problems/rotate-list/
func rotateRight() {
	head := sliceToLinkedList([]int{1, 2, 3, 4, 5})
	k := 2
	resultLinkedList := Leetcode{}.rotateRight(head, k)
	result := linkedListToSlice(resultLinkedList)
	fmt.Println("Rotate list:", result)
}

// 62: /problems/unique-paths/
func uniquePaths() {
	m := 3
	n := 7
	result := Leetcode{}.uniquePaths(m, n)
	fmt.Println("Unique paths:", result)
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
	mergeKLists()
	swapPairs()
	reverseKGroup()
	removeDuplicates()
	removeElement()
	strStr()
	divide()
	findSubstring()
	nextPermutation()
	longestValidParentheses()
	search()
	searchRange()
	searchInsert()
	isValidSudoku()
	solveSudoku()
	countAndSay()
	combinationSum()
	combinationSum2()
	firstMissingPositive()
	trap()
	multiply()
	isMatchWildcard()
	jump()
	permute()
	permuteUnique()
	rotate()
	groupAnagrams()
	myPow()
	solveNQueens()
	totalNQueens()
	maxSubArray()
	spiralOrder()
	canJump()
	mergeIntervals()
	insertInterval()
	lengthOfLastWord()
	generateMatrix()
	getPermutation()
	rotateRight()
	uniquePaths()
}
