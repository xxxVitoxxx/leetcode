package main

import "sort"

// time complexity: O(n * log n)
// space complexity: O(n)
func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}

	freq, index := make([][]int, len(count)), 0
	for k, v := range count {
		freq[index] = []int{v, k}
		index++
	}

	sort.Slice(freq, func(i, j int) bool {
		return freq[i][0] > freq[j][0]
	})

	answer := make([]int, k)
	for i := 0; i < k; i++ {
		answer[i] = freq[i][1]
	}

	return answer
}
