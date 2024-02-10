package main

// 1: /problems/two-sum/
func twoSum(nums []int, target int) []int {
	dic := make(map[int]int)
	for i, n := range nums {
		if idx, ok := dic[n]; ok {
			return []int{idx, i}
		}
		dic[target-n] = i
	}
	return nil
}
