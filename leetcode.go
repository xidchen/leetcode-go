package main

import (
	"math"
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
