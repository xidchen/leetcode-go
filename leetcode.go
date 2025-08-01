package main

import (
	"container/heap"
	"math"
	"sort"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type ListNodeHeap []*ListNode

func (h *ListNodeHeap) Len() int           { return len(*h) }
func (h *ListNodeHeap) Less(i, j int) bool { return (*h)[i].Val < (*h)[j].Val }
func (h *ListNodeHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *ListNodeHeap) Push(x interface{}) { *h = append(*h, x.(*ListNode)) }
func (h *ListNodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func sliceToLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	dummy := &ListNode{}
	current := dummy
	for _, num := range nums {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}
	return dummy.Next
}

func linkedListToSlice(node *ListNode) []int {
	var nums []int
	for node != nil {
		nums = append(nums, node.Val)
		node = node.Next
	}
	return nums
}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func reverseListElements(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func contains(arr []byte, val byte) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

type Leetcode struct{}

// 1: /problems/two-sum/
func (l Leetcode) twoSum(nums []int, target int) []int {
	dic := make(map[int]int)
	for i, n := range nums {
		if idx, ok := dic[n]; ok {
			return []int{idx, i}
		}
		dic[target-n] = i
	}
	return nil
}

// 2: /problems/add-two-numbers/
func (l Leetcode) addTwoNumbers(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	carry := 0
	for l1 != nil || l2 != nil || carry > 0 {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}
		current.Next = &ListNode{Val: carry % 10}
		current = current.Next
		carry /= 10
	}
	return dummy.Next
}

// 3: /problems/longest-substring-without-repeating-characters/
func (l Leetcode) lengthOfLongestSubstring(s string) int {
	lastSeen := make(map[rune]int)
	start, longest := 0, 0
	for i, c := range s {
		if lastIndex, exists := lastSeen[c]; exists {
			if lastIndex >= start {
				start = lastIndex + 1
			}
		}
		if i-start+1 > longest {
			longest = i - start + 1
		}
		lastSeen[c] = i
	}
	return longest
}

// 5: /problems/longest-palindromic-substring/
func (l Leetcode) longestPalindromeSubstring(s string) string {
	if len(s) < 2 || _isPalindromicString(s) {
		return s
	}
	var start, ml = -1, 0
	for i := 0; i < len(s); i++ {
		var odd, even string
		if i-ml-1 >= 0 {
			odd = s[i-ml-1 : i+1]
		}
		if i-ml >= 0 {
			even = s[i-ml : i+1]
		}
		if len(odd) > 0 && _isPalindromicString(odd) {
			start = i - ml - 1
			ml += 2
			continue
		}
		if len(even) > 0 && _isPalindromicString(even) {
			start = i - ml
			ml += 1
		}
	}
	return s[start : start+ml]
}

func _isPalindromicString(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

// 6: /problems/zigzag-conversion/
func (l Leetcode) convert(s string, numRows int) string {
	if numRows == 1 || len(s) < numRows {
		return s
	}
	zigzag := make([]string, numRows)
	row := 0
	step := 1
	for _, c := range s {
		zigzag[row] += string(c)
		if row == 0 {
			step = 1
		}
		if row == numRows-1 {
			step = -1
		}
		row += step
	}
	var res string
	for _, str := range zigzag {
		res += str
	}
	return res
}

// 7: /problems/reverse-integer/
func (l Leetcode) reverse(x int) int {
	negative := x < 0
	if negative {
		x = -x
	}
	var y = 0
	for x != 0 {
		y = y*10 + x%10
		x /= 10
	}
	if y > math.MaxInt32 {
		return 0
	}
	if negative {
		return -y
	} else {
		return y
	}

}

// 8: /problems/string-to-integer-atoi/
func (l Leetcode) myAtoi(s string) int {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0
	}
	var sign int64 = 1
	if s[0] == '+' || s[0] == '-' {
		if s[0] == '-' {
			sign = -1
		}
		s = s[1:]
	}
	res := int64(0)
	for _, c := range s {
		if c < '0' || c > '9' {
			break
		}
		res = res*10 + int64(c-'0')
		if res*sign > math.MaxInt32 {
			return math.MaxInt32
		} else if res*sign < math.MinInt32 {
			return math.MinInt32
		}
	}
	return int(res * sign)
}

// 9: /problems/palindrome-number/
func (l Leetcode) isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	return s == reverseString(s)
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// 10: /problems/regular-expression-matching/
func (l Leetcode) isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	dp[m][n] = true
	for i := m; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			firstMatch := i < m && (p[j] == s[i] || p[j] == '.')
			if j+1 < n && p[j+1] == '*' {
				dp[i][j] = dp[i][j+2] || (firstMatch && dp[i+1][j])
			} else {
				dp[i][j] = firstMatch && dp[i+1][j+1]
			}
		}
	}
	return dp[0][0]
}

// 11: /problems/container-with-most-water/
func (l Leetcode) maxArea(height []int) int {
	maxArea, i, j := 0, 0, len(height)-1
	for i < j {
		maxArea = max(maxArea, min(height[i], height[j])*(j-i))
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return maxArea
}

// 12: /problems/integer-to-roman/
func (l Leetcode) intToRoman(num int) string {
	mapping := [][]interface{}{
		{1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"},
		{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
		{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"},
		{1, "I"},
	}
	romans := strings.Builder{}
	for _, pair := range mapping {
		for pair[0].(int) <= num {
			num -= pair[0].(int)
			romans.WriteString(pair[1].(string))
		}
	}
	return romans.String()
}

// 13: /problems/roman-to-integer/
func (l Leetcode) romanToInt(s string) int {
	d := map[rune]int{
		'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000,
	}
	integer, prevInt := 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		roman := rune(s[i])
		if d[roman] >= prevInt {
			prevInt = d[roman]
			integer += d[roman]
		} else {
			integer -= d[roman]
		}
	}
	return integer
}

// 14: /problems/longest-common-prefix/
func (l Leetcode) longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	sort.Strings(strs)
	var res string
	for i := 0; i < len(strs[0]) && i < len(strs[len(strs)-1]); i++ {
		if strs[0][i] == strs[len(strs)-1][i] {
			res += string(strs[0][i])
		} else {
			break
		}
	}
	return res
}

// 15: /problems/3sum
func (l Leetcode) threeSum(nums []int) [][]int {
	dic := make(map[int]int)
	var res [][]int
	for _, n := range nums {
		dic[n]++
	}
	var sortedNums []int
	for n := range dic {
		sortedNums = append(sortedNums, n)
	}
	sort.Ints(sortedNums)
	for i, x := range sortedNums {
		if x == 0 && dic[x] > 2 {
			res = append(res, []int{0, 0, 0})
		} else if x != 0 && dic[x] > 1 && dic[-2*x] > 0 {
			res = append(res, []int{x, x, -2 * x})
		}
		if x < 0 {
			left := sort.Search(len(sortedNums)-i-1, func(j int) bool {
				return sortedNums[i+j+1] >= -x-sortedNums[len(sortedNums)-1]
			})
			right := sort.Search(len(sortedNums)-i-1, func(j int) bool {
				return sortedNums[i+j+1] > x/-2
			})
			for _, y := range sortedNums[i+left+1 : i+right+1] {
				z := -x - y
				if dic[z] > 0 && z != y {
					res = append(res, []int{x, y, z})
				}
			}
		}
	}
	return res
}

// 16: /problems/3sum-closest/
func (l Leetcode) threeSumClosest(nums []int, target int) int {
	n := len(nums)
	sort.Ints(nums)
	res := nums[0] + nums[1] + nums[2]
	for i := range n - 2 {
		j, k := i+1, n-1
		if nums[i]+nums[j]+nums[j+1] >= target {
			k = j + 1
		}
		if nums[i]+nums[k-1]+nums[k] <= target {
			j = k - 1
		}
		for j < k {
			s := nums[i] + nums[j] + nums[k]
			if _abs(target-s) < _abs(target-res) {
				res = s
			}
			if s == target {
				return res
			}
			if s < target {
				j++
			}
			if s > target {
				k--
			}
		}
	}
	return res
}

func _abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 17: /problems/letter-combinations-of-a-phone-number/
func (l Leetcode) letterCombinations(digits string) []string {
	var res []string
	if len(digits) == 0 || _contains(digits, '0') || _contains(digits, '1') {
		return res
	}
	mapping := map[rune][]rune{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}
	res = append(res, "")
	for _, digit := range digits {
		var temp []string
		for _, result := range res {
			for _, letter := range mapping[digit] {
				temp = append(temp, result+string(letter))
			}
		}
		res = temp
	}
	return res
}

func _contains(s string, r rune) bool {
	for _, c := range s {
		if c == r {
			return true
		}
	}
	return false
}

// 18: /problems/4sum/
func (l Leetcode) fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	return _kSum(nums, target, 4)
}

func _kSum(n []int, t int, k int) [][]int {
	var res [][]int
	if len(n) < k || t < n[0]*k || n[len(n)-1]*k < t {
		return res
	}
	if k == 2 {
		return _twoSum(n, t)
	}
	for i := 0; i < len(n); i++ {
		if i == 0 || n[i-1] != n[i] {
			sets := _kSum(n[i+1:], t-n[i], k-1)
			for _, set := range sets {
				res = append(res, append([]int{n[i]}, set...))
			}
		}
	}
	return res
}

func _twoSum(n []int, t int) [][]int {
	var res [][]int
	lo, hi := 0, len(n)-1
	for lo < hi {
		sum := n[lo] + n[hi]
		if sum < t || (lo > 0 && n[lo] == n[lo-1]) {
			lo++
		} else if sum > t || (hi < len(n)-1 && n[hi] == n[hi+1]) {
			hi--
		} else {
			res = append(res, []int{n[lo], n[hi]})
			lo++
			hi--
		}
	}
	return res
}

// 19: /problems/remove-nth-node-from-end-of-list/
func (l Leetcode) removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	first := dummy
	second := dummy
	for i := 0; i < n; i++ {
		if first.Next != nil {
			first = first.Next
		} else {
			return head
		}
	}
	for first.Next != nil {
		first = first.Next
		second = second.Next
	}
	if second.Next != nil {
		second.Next = second.Next.Next
	}
	return dummy.Next
}

// 20: /problems/valid-parentheses/
func (l Leetcode) isValid(s string) bool {
	d := map[rune]rune{'(': ')', '[': ']', '{': '}'}
	var stack []rune
	for _, c := range s {
		if closing, exists := d[c]; exists {
			stack = append(stack, closing)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != c {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

// 21: /problems/merge-two-sorted-lists/
func (l Leetcode) mergeTwoLists(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	prev := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			prev.Next = l1
			l1 = l1.Next
		} else {
			prev.Next = l2
			l2 = l2.Next
		}
		prev = prev.Next
	}
	if l1 != nil {
		prev.Next = l1
	} else {
		prev.Next = l2
	}
	return dummy.Next
}

// 22: /problems/generate-parentheses/
func (l Leetcode) generateParenthesis(n int) []string {
	var res []string
	var backtrack func(s string, left int, right int)

	backtrack = func(s string, left int, right int) {
		if len(s) == 2*n {
			res = append(res, s)
			return
		}
		if left < n {
			backtrack(s+"(", left+1, right)
		}
		if right < left {
			backtrack(s+")", left, right+1)
		}
	}

	backtrack("", 0, 0)
	return res
}

// 23: /problems/merge-k-sorted-lists/
func (l Leetcode) mergeKLists(lists []*ListNode) *ListNode {
	minHeap := &ListNodeHeap{}
	heap.Init(minHeap)
	for _, node := range lists {
		if node != nil {
			heap.Push(minHeap, node)
		}
	}
	dummy := &ListNode{}
	prev := dummy
	for minHeap.Len() > 0 {
		node := heap.Pop(minHeap).(*ListNode)
		prev.Next = node
		prev = prev.Next
		if node.Next != nil {
			heap.Push(minHeap, node.Next)
		}
	}
	return dummy.Next
}

// 24: /problems/swap-nodes-in-pairs/
func (l Leetcode) swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy
	for prev.Next != nil && prev.Next.Next != nil {
		first := prev.Next
		second := first.Next
		first.Next = second.Next
		second.Next = first
		prev.Next = second
		prev = first
	}
	return dummy.Next
}

// 25: /problems/reverse-nodes-in-k-group/
func (l Leetcode) reverseKGroup(head *ListNode, k int) *ListNode {
	if k < 2 {
		return head
	}
	node := head
	for i := 0; i < k; i++ {
		if node == nil {
			return head
		}
		node = node.Next
	}
	prev := l.reverseKGroup(node, k)
	for i := 0; i < k; i++ {
		temp := head.Next
		head.Next = prev
		prev = head
		head = temp
	}
	return prev
}

// 26: /problems/remove-duplicates-from-sorted-array/
func (l Leetcode) removeDuplicates(nums []int) int {
	nextNew := 0
	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i] != nums[i-1] {
			nums[nextNew] = nums[i]
			nextNew++
		}
	}
	return nextNew
}

// 27: /problems/remove-element/
func (l Leetcode) removeElement(nums []int, val int) int {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}

// 28: /problems/find-the-index-of-the-first-occurrence-in-a-string/
func (l Leetcode) strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

// 29: /problems/divide-two-integers/
func (l Leetcode) divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	diffSign := (dividend < 0) != (divisor < 0)
	dividend64 := int64(dividend)
	divisor64 := int64(divisor)
	res := 0
	dividend64 = abs(dividend64)
	divisor64 = abs(divisor64)
	for dividend64 >= divisor64 {
		temp := divisor64
		multiple := 1
		for dividend64 >= temp<<1 {
			temp <<= 1
			multiple <<= 1
		}
		dividend64 -= temp
		res += multiple
	}
	if diffSign {
		res = -res
	}
	return max(min(res, math.MaxInt32), math.MinInt32)
}

// 30: /problems/substring-with-concatenation-of-all-words/
func (l Leetcode) findSubstring(s string, words []string) []int {
	var res []int
	if len(words) == 0 || len(s) < len(words)*len(words[0]) {
		return res
	}
	wc := len(words)
	wl := len(words[0])
	sl := len(s)
	wd := make(map[string]int)
	for _, w := range words {
		wd[w]++
	}
	for i := 0; i < wl; i++ {
		start := i
		cnt := 0
		tmpDict := make(map[string]int)
		for j := i; j <= sl-wl; j += wl {
			word := s[j : j+wl]
			if count, exists := wd[word]; exists {
				cnt++
				tmpDict[word]++
				for tmpDict[word] > count {
					startWord := s[start : start+wl]
					tmpDict[startWord]--
					start += wl
					cnt--
				}
				if cnt == wc {
					res = append(res, start)
					startWord := s[start : start+wl]
					tmpDict[startWord]--
					start += wl
					cnt--
				}
			} else {
				start = j + wl
				cnt = 0
				tmpDict = make(map[string]int)
			}
		}
	}
	return res
}

// 31: /problems/next-permutation/
func (l Leetcode) nextPermutation(nums []int) {
	n := len(nums)
	if n < 2 {
		return
	}
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		j := n - 1
		for j >= 0 && nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	reverseListElements(nums[i+1:])
}

// 32: /problems/longest-valid-parentheses/
func (l Leetcode) longestValidParentheses(s string) int {
	stack := []int{-1}
	maxLen := 0
	for i, ch := range s {
		if ch == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				maxLen = max(maxLen, i-stack[len(stack)-1])
			}
		}
	}
	return maxLen
}

// 33: /problems/search-in-rotated-sorted-array/
func (l Leetcode) search(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	left, right := 0, n-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[left] <= nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

// 34: /problems/find-first-and-last-position-of-element-in-sorted-array/
func (l Leetcode) searchRange(nums []int, target int) []int {
	binary := func(tgt float64, left, right int) int {
		for left <= right {
			mid := (left + right) / 2
			if float64(nums[mid]) < tgt {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		return left
	}
	n := len(nums)
	lower := binary(float64(target)-0.5, 0, n-1)
	upper := binary(float64(target)+0.5, 0, n-1)
	if lower == upper {
		return []int{-1, -1}
	}
	return []int{lower, upper - 1}
}

// 35: /problems/search-insert-position/
func (l Leetcode) searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

// 36: /problems/valid-sudoku/
func (l Leetcode) isValidSudoku(board [][]byte) bool {
	rows := make([][]byte, 9)
	cols := make([][]byte, 9)
	boxes := make([][]byte, 9)
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			digit := board[r][c]
			if digit == '.' {
				continue
			}
			b := r/3*3 + c/3
			if contains(rows[r], digit) ||
				contains(cols[c], digit) ||
				contains(boxes[b], digit) {
				return false
			}
			rows[r] = append(rows[r], digit)
			cols[c] = append(cols[c], digit)
			boxes[b] = append(boxes[b], digit)
		}
	}
	return true
}

// 37: /problems/sudoku-solver/
func (l Leetcode) solveSudoku(board [][]byte) {
	backtrackSudoku(board)
}

func backtrackSudoku(board [][]byte) bool {
	row, col, found := findEmptySudokuCell(board)
	if !found {
		return true
	}
	for digit := '1'; digit <= '9'; digit++ {
		if isValidSudokuPlacement(board, row, col, byte(digit)) {
			board[row][col] = byte(digit)
			if backtrackSudoku(board) {
				return true
			}
			board[row][col] = '.'
		}
	}
	return false
}

func findEmptySudokuCell(board [][]byte) (int, int, bool) {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == '.' {
				return row, col, true
			}
		}
	}
	return 0, 0, false
}

func isValidSudokuPlacement(board [][]byte, row, col int, digit byte) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == digit {
			return false
		}
	}
	for i := 0; i < 9; i++ {
		if board[i][col] == digit {
			return false
		}
	}
	boxRow, boxCol := (row/3)*3, (col/3)*3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[boxRow+i][boxCol+j] == digit {
				return false
			}
		}
	}
	return true
}

// 38: /problems/count-and-say/
func (l Leetcode) countAndSay(n int) string {
	seq := []int{1}
	for i := 1; i < n; i++ {
		var next []int
		for _, num := range seq {
			if len(next) == 0 || num != next[len(next)-1] {
				next = append(next, 1)
				next = append(next, num)
			} else {
				next[len(next)-2]++
			}
		}
		seq = next
	}
	var sb strings.Builder
	for _, num := range seq {
		sb.WriteString(strconv.Itoa(num))
	}
	return sb.String()
}

// 39: /problems/combination-sum/
func (l Leetcode) combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var helper func(nums []int, nxt int, t int, p []int, r *[][]int)
	helper = func(nums []int, nxt int, t int, p []int, r *[][]int) {
		if t == 0 {
			pCopy := make([]int, len(p))
			copy(pCopy, p)
			*r = append(*r, pCopy)
			return
		}
		if nxt == len(nums) {
			return
		}
		for i := 0; t-i*nums[nxt] >= 0; i++ {
			currentP := make([]int, len(p))
			copy(currentP, p)
			for j := 0; j < i; j++ {
				currentP = append(currentP, nums[nxt])
			}
			helper(nums, nxt+1, t-i*nums[nxt], currentP, r)
		}
	}
	helper(candidates, 0, target, []int{}, &res)
	return res
}

// 40: /problems/combination-sum-ii/
func (l Leetcode) combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	var current []int
	sort.Ints(candidates)
	var backtrack func(pos, remain int)
	backtrack = func(pos, remain int) {
		if remain == 0 {
			pCopy := make([]int, len(current))
			copy(pCopy, current)
			res = append(res, pCopy)
			return
		}
		for i := pos; i < len(candidates); i++ {
			if i > pos && candidates[i] == candidates[i-1] {
				continue
			}
			if candidates[i] > remain {
				break
			}
			current = append(current, candidates[i])
			backtrack(i+1, remain-candidates[i])
			current = current[:len(current)-1]
		}
	}
	backtrack(0, target)
	return res
}

// 41: /problems/first-missing-positive/
func (l Leetcode) firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}

// 42: /problems/trapping-rain-water/
func (l Leetcode) trap(height []int) int {
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	res := 0
	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				res += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				res += rightMax - height[right]
			}
			right--
		}
	}
	return res
}

// 43: /problems/multiply-strings/
func (l Leetcode) multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	m, n := len(num1), len(num2)
	res := make([]int, m+n)
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			mul := int(num1[i]-'0') * int(num2[j]-'0')
			p1 := i + j
			p2 := i + j + 1
			sum := mul + res[p2]
			res[p1] += sum / 10
			res[p2] = sum % 10
		}
	}
	var sb strings.Builder
	for i := 0; i < len(res); i++ {
		if !(sb.Len() == 0 && res[i] == 0) {
			sb.WriteString(strconv.Itoa(res[i]))
		}
	}
	return sb.String()
}

// 44: /problems/wildcard-matching/
func (l Leetcode) isMatchWildcard(s string, p string) bool {
	sIdx, pIdx := 0, 0
	starIdx, match := -1, 0
	for sIdx < len(s) {
		if pIdx < len(p) && (p[pIdx] == '?' || p[pIdx] == s[sIdx]) {
			sIdx++
			pIdx++
		} else if pIdx < len(p) && p[pIdx] == '*' {
			starIdx = pIdx
			match = sIdx
			pIdx++
		} else if starIdx != -1 {
			pIdx = starIdx + 1
			match++
			sIdx = match
		} else {
			return false
		}
	}
	for pIdx < len(p) && p[pIdx] == '*' {
		pIdx++
	}
	return pIdx == len(p)
}

// 45: /problems/jump-game-ii/
func (l Leetcode) jump(nums []int) int {
	jumps := 0
	currentEnd := 0
	farthest := 0
	for i := 0; i < len(nums)-1; i++ {
		farthest = max(farthest, i+nums[i])
		if i == currentEnd {
			jumps++
			currentEnd = farthest
		}
	}
	return jumps
}

// 46: /problems/permutations/
func (l Leetcode) permute(nums []int) [][]int {
	var res [][]int
	var backtrack func(start int)
	backtrack = func(start int) {
		if start == len(nums) {
			perm := make([]int, len(nums))
			copy(perm, nums)
			res = append(res, perm)
			return
		}
		for i := start; i < len(nums); i++ {
			nums[start], nums[i] = nums[i], nums[start]
			backtrack(start + 1)
			nums[start], nums[i] = nums[i], nums[start]
		}
	}
	backtrack(0)
	return res
}

// 47: /problems/permutations-ii/
func (l Leetcode) permuteUnique(nums []int) [][]int {
	var res [][]int
	var backtrack func(start int)
	backtrack = func(start int) {
		used := make(map[int]bool)
		if start == len(nums) {
			perm := make([]int, len(nums))
			copy(perm, nums)
			res = append(res, perm)
			return
		}
		for i := start; i < len(nums); i++ {
			if used[nums[i]] {
				continue
			}
			used[nums[i]] = true
			nums[start], nums[i] = nums[i], nums[start]
			backtrack(start + 1)
			nums[start], nums[i] = nums[i], nums[start]
		}
	}
	backtrack(0)
	return res
}

// 48: /problems/rotate-image/
func (l Leetcode) rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	for i := 0; i < n; i++ {
		reverseListElements(matrix[i])
	}
}

// 49: /problems/group-anagrams/
func (l Leetcode) groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)
	for _, str := range strs {
		count := make([]byte, 26)
		for i := 0; i < len(str); i++ {
			count[str[i]-'a']++
		}
		key := string(count)
		groups[key] = append(groups[key], str)
	}
	res := make([][]string, 0, len(groups))
	for _, group := range groups {
		res = append(res, group)
	}
	return res
}
