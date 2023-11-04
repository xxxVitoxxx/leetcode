package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	length := len(nums)
	if length < 4 {
		return [][]int{}
	}

	var res [][]int
	for i := 0; i < length-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		leftSum := nums[i] + nums[i+1] + nums[i+2] + nums[i+3]
		if leftSum > target {
			break
		}

		rightSum := nums[i] + nums[length-1] + nums[length-2] + nums[length-3]
		if rightSum < target {
			continue
		}

		for j := i + 1; j < length-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			k, l := j+1, length-1
			for k < l {
				num := nums[i] + nums[j] + nums[k] + nums[l]
				if num == target {
					res = append(res, []int{nums[i], nums[j], nums[k], nums[l]})
					k, l = k+1, l-1
					for k < l && nums[k] == nums[k-1] {
						k++
					}
					for k < l && nums[l] == nums[l+1] {
						l--
					}
				} else if num > target {
					l--
				} else {
					k++
				}
			}
		}
	}
	return res
}
