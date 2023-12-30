package main

func longestOnes(nums []int, k int) int {
	ans, left, right, zero := 0, 0, 0, 0
	for right < len(nums) {
		if nums[right] == 1 {
			right++
		} else {
			if zero < k {
				zero++
				right++
			} else {
				for zero == k {
					if nums[left] == 0 {
						zero--
					}
					left++
				}
			}
		}
		if right-left > ans {
			ans = right - left
		}
	}
	return ans
}

// optimize longestOnes
// time complexity: O(n)
// space complexity: O(1)
func longestOnes2(nums []int, k int) int {
	ans, zero := 0, 0
	for left, right := 0, 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zero++
		}
		for zero > k {
			if nums[left] == 0 {
				zero--
			}
			left++
		}
		if right-left+1 > ans {
			ans = right - left + 1
		}
	}
	return ans
}

func longestOnes3(nums []int, k int) int {
	left, right := 0, 0
	for ; right < len(nums); right++ {
		if nums[right] == 0 {
			k--
		}

		if k < 0 {
			// there are more zeros in our window than k
			if nums[left] == 0 {
				k++
			}
			left++
		}
	}
	return right - left
}
