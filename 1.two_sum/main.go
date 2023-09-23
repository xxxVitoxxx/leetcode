package main

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, curr := range nums {
		if val, ok := m[target-curr]; ok {
			return []int{val, i}
		}
		m[curr] = i
	}
	return nil
}
