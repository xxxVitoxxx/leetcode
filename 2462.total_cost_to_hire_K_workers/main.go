package main

import "container/heap"

// heap (priority queue)
// c is the given integer candidates.
// time complexity: O((c+k) * log c)
//   - we need to initialize one priority queue of size up to 2 * c, which takes O(c * log c) time.
//   - during k hiring rounds, we keep popping top elements from priority queue and pushing
//     new elements into priority queue for up to k times. operations on a priority queue take
//     amortized O(log c) time. thus this process takes O(k * log c) time.
//   - therefore, the time complexity is O((c+k) * log c)
//
// space complexity: O(c)
//   - we need to store at most 2 * c elements(the first c and the last c elements) of costs in
//     the priority queue.
func totalCost(costs []int, k int, candidates int) int64 {
	minHeap := &MinHeap{}

	// add first candidates to heap
	for i := 0; i < candidates; i++ {
		heap.Push(minHeap, []int{costs[i], 0})
	}

	// add last candidates to heap
	for i := max(candidates, len(costs)-candidates); i < len(costs); i++ {
		heap.Push(minHeap, []int{costs[i], 1})
	}

	left := candidates
	right := len(costs) - 1 - candidates
	answer := 0
	for i := 0; i < k; i++ {
		candidate := heap.Pop(minHeap).([]int)
		answer += candidate[0]
		index := candidate[1]

		if left <= right {
			if index == 0 {
				heap.Push(minHeap, []int{costs[left], 0})
				left++
			} else {
				heap.Push(minHeap, []int{costs[right], 1})
				right--
			}
		}
	}

	return int64(answer)
}

type MinHeap [][]int

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool {
	if h[i][0] == h[j][0] {
		return h[i][1] < h[j][1]
	}
	return h[i][0] < h[j][0]
}
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}
func (h *MinHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
