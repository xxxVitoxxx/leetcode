package main

// brute force
func sortArrayByParityII(nums []int) []int {
	even, odd := []int{}, []int{}
	for _, num := range nums {
		if num%2 == 0 {
			even = append(even, num)
		} else {
			odd = append(odd, num)
		}
	}

	res := make([]int, len(nums))
	for i := 0; i < len(res); i++ {
		if i%2 == 0 {
			res[i], even = even[0], even[1:]
		} else {
			res[i], odd = odd[0], odd[1:]
		}
	}
	return res
}

func sortArrayByParityII2(nums []int) []int {
	res := make([]int, len(nums))
	evenPos, oddPos := 0, 0
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			for evenPos < len(nums) && nums[evenPos]%2 == 1 {
				evenPos++
			}
			res[i] = nums[evenPos]
			evenPos++
		} else {
			for oddPos < len(nums) && nums[oddPos]%2 == 0 {
				oddPos++
			}
			res[i] = nums[oddPos]
			oddPos++
		}
	}
	return res
}

// 記錄有問題的 index 並只對有問題的 index 做處理
func sortArrayByParityII3(nums []int) []int {
	even, odd, index := []int{}, []int{}, []int{}
	for i := range nums {
		if i%2 == 0 && nums[i]%2 == 1 {
			odd = append(odd, nums[i])
			index = append(index, i)
		} else if i%2 == 1 && nums[i]%2 == 0 {
			even = append(even, nums[i])
			index = append(index, i)
		}
	}

	for _, idx := range index {
		if idx%2 == 0 {
			nums[idx], even = even[0], even[1:]
		} else if idx%2 == 1 {
			nums[idx], odd = odd[0], odd[1:]
		}
	}
	return nums
}

// use in-place
func sortArrayByParityII4(nums []int) []int {
	even, odd := 0, 1
	for even < len(nums) && odd < len(nums) {
		if nums[even]%2 == 0 {
			even += 2
		} else if nums[odd]%2 == 1 {
			odd += 2
		} else {
			nums[even], nums[odd] = nums[odd], nums[even]
			even, odd = even+2, odd+2
		}
	}
	return nums
}
