package main

// brute force
func duplicateZeros(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == 0 {
			i++
			tmp := arr[i]
			arr[i] = 0
			for j := i + 1; j < len(arr); j++ {
				arr[j], tmp = tmp, arr[j]
			}
		}
	}
}

func duplicateZeros2(arr []int) {
	var nums []int
	for _, num := range arr {
		if num > 0 {
			nums = append(nums, num)
		} else {
			nums = append(nums, num, 0)
		}
	}

	for i := 0; i < len(arr); i++ {
		arr[i] = nums[i]
	}
}

// use copy
func duplicateZeros3(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == 0 {
			copy(arr[i+1:], arr[i:])
			i++
		}
	}
}

func duplicateZeros4(arr []int) {
	l, zeroCount := len(arr), 0
	for _, num := range arr {
		if num == 0 {
			zeroCount++
		}
	}

	for i := l - 1; i >= 0; i-- {
		if arr[i] == 0 {
			zeroCount--
			// duplicate zero
			if i+zeroCount+1 < l {
				arr[i+zeroCount+1] = 0
			}

			if i+zeroCount < l {
				arr[i+zeroCount] = 0
			}
		} else {
			if i+zeroCount < l {
				arr[i+zeroCount] = arr[i]
			}
		}
	}
}
