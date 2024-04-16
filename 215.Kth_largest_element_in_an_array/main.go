package main

import (
	"container/heap"
	"math"
	"math/rand"
)

// hash
// time complexity: O(d) where d is max - min + 1
// space complexity: O(n)
func findKthLargest(nums []int, k int) int {
	m := make(map[int]int)
	max, min := math.MinInt, math.MaxInt
	for _, num := range nums {
		m[num]++
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}

	for i := max; i >= min; i-- {
		frequency, ok := m[i]
		if ok {
			k -= frequency
			if k <= 0 {
				return i
			}
		}
	}

	return -1
}

// heap
// time complexity: O(n + k * logn) building a heap takes O(n) and removing an element from the heap takes O(logn) time.
// space complexity: O(n) store the numbers in the heap.
func findKthLargest2(nums []int, k int) int {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		heapify(nums, i, len(nums)-1)
	}

	answer, lastIndex := 0, len(nums)-1
	for i := 0; i < k; i++ {
		answer = nums[0]
		nums[0], nums[lastIndex] = nums[lastIndex], nums[0]
		heapify(nums, 0, lastIndex-1)
		lastIndex--
	}

	return answer
}

func heapify(nums []int, parent, lastIndex int) {
	largest := parent
	left, right := 2*parent+1, 2*parent+2
	if left <= lastIndex && nums[left] > nums[largest] {
		largest = left
	}
	if right <= lastIndex && nums[right] > nums[largest] {
		largest = right
	}
	if largest != parent {
		nums[parent], nums[largest] = nums[largest], nums[parent]
		heapify(nums, largest, lastIndex)
	}
}

// heap
// time complexity: O(n * logk) each of the n elements is processed once. however, heap operations take O(logk) time,
// leading to an overall complexity of O(n * logk).
// space complexity: O(k) use a heap with a maximum of k elements.
func findKthLargest3(nums []int, k int) int {
	intHeap := IntHeap(nums[:k])
	heap.Init(&intHeap)

	for _, num := range nums[k:] {
		if num > intHeap[0] {
			heap.Pop(&intHeap)
			heap.Push(&intHeap, num)
		}
	}

	return intHeap[0]
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// quickly select
// time complexity: O(n) in the worst case is O(n * n)
// space complexity: O(n) we need n space to create left, mid, and right.
// other implementations of quickly select can avoid creating these three in memory,
// but in the worst case, those implementations would still require O(n) space for recursion call stack.
func findKthLargest4(nums []int, k int) int {
	pivot := nums[rand.Intn(len(nums))]
	var left, mid, right []int
	for _, num := range nums {
		if num > pivot {
			left = append(left, num)
		} else if num < pivot {
			right = append(right, num)
		} else {
			mid = append(mid, num)
		}
	}

	if k <= len(left) {
		return findKthLargest3(left, k)
	}
	if len(left)+len(mid) < k {
		return findKthLargest3(right, k-len(left)-len(mid))
	}

	return pivot
}

// quickly select
// time complexity: O(n) in the worst case is O(n * n)
// space complexity: O(n)
func findKthLargest5(nums []int, k int) int {
	return quicklySelect(nums, k, 0, len(nums)-1)
}

func quicklySelect(nums []int, k, left, right int) int {
	pivot, p := nums[right], left
	for i := left; i < right; i++ {
		if nums[i] < pivot {
			nums[i], nums[p] = nums[p], nums[i]
			p++
		}
	}
	nums[p], nums[right] = nums[right], nums[p]

	if len(nums)-k < p {
		return quicklySelect(nums, k, left, p-1)
	} else if len(nums)-k > p {
		return quicklySelect(nums, k, p+1, right)
	}

	return nums[p]
}
