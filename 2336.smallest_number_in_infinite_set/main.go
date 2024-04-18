package main

import "container/heap"

// hash
type SmallestInfiniteSet struct {
	present map[int]bool
}

func Constructor() SmallestInfiniteSet {
	present := make(map[int]bool, 1000)
	for i := 1; i <= 1000; i++ {
		present[i] = true
	}
	return SmallestInfiniteSet{present}
}

func (s *SmallestInfiniteSet) PopSmallest() int {
	var answer int
	for i := 1; i <= 1000; i++ {
		if s.present[i] {
			s.present[i] = false
			answer = i
			break
		}
	}
	return answer
}

func (s *SmallestInfiniteSet) AddBack(num int) {
	s.present[num] = true
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	l := len(old)
	x := old[l-1]
	*h = old[:l-1]
	return x
}

type SmallestInfiniteSet2 struct {
	smallest int
	minHeap  *IntHeap
	heapSet  map[int]bool
}

func Constructor2() SmallestInfiniteSet2 {
	intHeap := new(IntHeap)
	heap.Init(intHeap)
	return SmallestInfiniteSet2{
		smallest: 1,
		minHeap:  intHeap,
		heapSet:  make(map[int]bool),
	}
}

func (s *SmallestInfiniteSet2) PopSmallest() int {
	var oldSmallest int
	if s.minHeap.Len() > 0 {
		oldSmallest = heap.Pop(s.minHeap).(int)
		delete(s.heapSet, oldSmallest)
	} else {
		oldSmallest = s.smallest
		s.smallest++
	}

	return oldSmallest
}

func (s *SmallestInfiniteSet2) AddBack(num int) {
	if ok := s.heapSet[num]; !ok && num < s.smallest {
		heap.Push(s.minHeap, num)
		s.heapSet[num] = true
	}
}

type SmallestInfiniteSet3 struct {
	removeItems map[int]bool
	smallest    int
}

func Constructor3() SmallestInfiniteSet3 {
	return SmallestInfiniteSet3{
		removeItems: make(map[int]bool),
		smallest:    1,
	}
}

func (s *SmallestInfiniteSet3) PopSmallest() int {
	answer := s.smallest
	s.removeItems[s.smallest] = true
	for s.removeItems[s.smallest] {
		s.smallest++
	}

	return answer
}

func (s *SmallestInfiniteSet3) AddBack(num int) {
	delete(s.removeItems, num)
	if num < s.smallest {
		s.smallest = num
	}
}

/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */
