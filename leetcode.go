package main

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

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
	if len(s) < 2 || isPalindromicString(s) {
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
		if len(odd) > 0 && isPalindromicString(odd) {
			start = i - ml - 1
			ml += 2
			continue
		}
		if len(even) > 0 && isPalindromicString(even) {
			start = i - ml
			ml += 1
		}
	}
	return s[start : start+ml]
}

func isPalindromicString(s string) bool {
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
	var result string
	for i := 0; i < len(strs[0]) && i < len(strs[len(strs)-1]); i++ {
		if strs[0][i] == strs[len(strs)-1][i] {
			result += string(strs[0][i])
		} else {
			break
		}
	}
	return result
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
