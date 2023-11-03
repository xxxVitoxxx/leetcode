package main

import "sort"

// use hash
type TwoSum struct {
	m map[int]int
}

func Constructor() TwoSum {
	return TwoSum{make(map[int]int)}
}

func (this *TwoSum) Add(number int) {
	this.m[number]++
}

func (this *TwoSum) Find(value int) bool {
	for k, v := range this.m {
		if _, ok := this.m[value-k]; ok && (value-k != k || v > 1) {
			return true
		}
	}
	return false
}

/**
 * Your TwoSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * param_2 := obj.Find(value);
 */

// use slice
type TwoSum2 struct {
	arr []int
}

func Constructor2() TwoSum2 {
	return TwoSum2{[]int{}}
}

func (this *TwoSum2) Add2(number int) {
	this.arr = append(this.arr, number)
	sort.Ints(this.arr)
}

func (this *TwoSum2) Find2(value int) bool {
	left, right := 0, len(this.arr)-1
	for left < right {
		num := this.arr[left] + this.arr[right]
		if num > value {
			right--
		} else if num < value {
			left++
		} else {
			return true
		}
	}
	return false
}

/**
 * Your TwoSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * param_2 := obj.Find(value);
 */
