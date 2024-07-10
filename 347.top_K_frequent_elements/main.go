package main

import (
	"container/heap"
	"math/rand"
	"sort"
)

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

// quick select
// time complexity: O(n) in the average case, O(n²) is the worst case.
// in the worst case of constantly badly chosen pivots, the problem is not divided
// by half at each step, it becomes just one element less, which leads to O(n²) time complexity.
// it happens, for example, if at each step you choose the pivot not randomly,
// but take the rightmost element. for the random pivot choice, the probability of
// having such a worst-case is negligibly small.
//
// space complexity: O(n) to store hash map and array of unique elements.
func topKFrequent2(nums []int, k int) []int {
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}

	unique := make([]int, 0, len(count))
	for k := range count {
		unique = append(unique, k)
	}

	partition := func(left, right, pivotIndex int) int {
		pivotFreq := count[unique[pivotIndex]]
		unique[pivotIndex], unique[right] = unique[right], unique[pivotIndex]

		storeIndex := left
		for i := left; i < right; i++ {
			if count[unique[i]] < pivotFreq {
				unique[i], unique[storeIndex] = unique[storeIndex], unique[i]
				storeIndex++
			}
		}

		unique[storeIndex], unique[right] = unique[right], unique[storeIndex]

		return storeIndex
	}

	var quickSelect func(int, int, int)
	quickSelect = func(left, right, kSmallest int) {
		if left == right {
			return
		}

		pivotIndex := rand.Intn(right-left+1) + left
		pivotIndex = partition(left, right, pivotIndex)
		if kSmallest == pivotIndex {
			return
		} else if kSmallest < pivotIndex {
			quickSelect(left, pivotIndex-1, kSmallest)
		} else {
			quickSelect(pivotIndex+1, right, kSmallest)
		}
	}

	n := len(unique)
	quickSelect(0, n-1, n-k)

	return unique[n-k:]
}

// time complexity: O(n)
// space complexity: O(n)
func topKFrequent3(nums []int, k int) []int {
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}

	// frequent element would be greater than or equal to 1
	// if just one number '1' appears in the nums and the length of nums is 2 that
	// represent freq[2] = 1, so the length of freq will be the length of nums + 1
	freq := make([][]int, len(nums)+1)
	for k, v := range count {
		freq[v] = append(freq[v], k)
	}

	answer := make([]int, 0, k)
	c := 0
	for i := len(freq) - 1; i >= 0; i-- {
		for _, num := range freq[i] {
			answer = append(answer, num)
			c++
			if c == k {
				return answer
			}
		}
	}

	return answer
}

// heap
// time complexity: O(n * log n)
// space complexity: O(n)
func topKFrequent4(nums []int, k int) []int {
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}

	minHeap := &MinHeap{}
	for num, freq := range count {
		heap.Push(minHeap, [2]int{num, freq})
		if minHeap.Len() > k {
			heap.Pop(minHeap)
		}
	}

	var answer []int
	for i := 0; i < k; i++ {
		answer = append(answer, heap.Pop(minHeap).([2]int)[0])
	}
	return answer
}

type MinHeap [][2]int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}
func (h *MinHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}
