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
