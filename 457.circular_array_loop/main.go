package main

// use fast slow pointers
func circularArrayLoop(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}

	for i := 0; i < len(nums); i++ {
		slow, fast := i, getNext(i, nums)
		for nums[slow]*nums[fast] > 0 {
			if slow == fast {
				if slow == getNext(slow, nums) {
					break
				}
				return true
			}

			if nums[fast]*nums[getNext(fast, nums)] < 0 {
				return false
			}

			slow = getNext(slow, nums)
			fast = getNext(getNext(fast, nums), nums)
		}

	}
	return false
}

func getNext(index int, nums []int) int {
	next := (index + nums[index]) % len(nums)
	if next >= 0 {
		return next
	}
	return next + len(nums)
}

func circularArrayLoop2(nums []int) bool {
	l, visited := len(nums), map[int]struct{}{}
	for i, num := range nums {
		if _, ok := visited[i]; ok {
			continue
		}

		prev, loop := i, map[int]struct{}{}
		for move, index := num, i; move != 0; move = nums[index] {
			index = (index + move%l + l) % l
			if num*nums[index] < 0 || prev == index {
				break
			}

			if _, ok := loop[index]; ok {
				return true
			}

			visited[index] = struct{}{}
			loop[index] = struct{}{}
			prev = index
		}
	}
	return false
}
