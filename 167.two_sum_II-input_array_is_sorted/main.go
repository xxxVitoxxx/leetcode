package main

// constraints:
// 1. numbers is sorted in non-decreasing order
// 2. must use only constant extra space

// brute force
// visit each number and calculate if the sum equals the target
func twoSum(numbers []int, target int) []int {
	res := make([]int, 2)
	j := 0
	for i := j + 1; i < len(numbers); i++ {
		sum := numbers[i] + numbers[j]
		if sum == target {
			res[0], res[1] = j+1, i+1
			break
		}

		if i == len(numbers)-1 {
			j++
			i = j
		}
	}
	return res
}

// use two pointer
func twoSum2(numbers []int, target int) []int {
	pnt1, pnt2 := 0, len(numbers)-1
	for pnt1 < pnt2 {
		sum := numbers[pnt1] + numbers[pnt2]
		if sum > target {
			pnt2--
		} else if sum < target {
			pnt1++
		} else {
			break
		}
	}
	return []int{pnt1 + 1, pnt2 + 1}
}
