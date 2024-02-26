package main

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
