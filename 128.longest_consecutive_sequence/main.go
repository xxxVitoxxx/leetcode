package main

// time complexity: O(N)
// space complexity: O(N)
func longestConsecutive(nums []int) int {
	m := make(map[int]bool)
	for _, num := range nums {
		m[num] = true
	}

	var result int
	for n := range m {
		if m[n-1] {
			continue
		}

		length := 1
		for {
			if !m[n+length] {
				break
			}

			length++
		}

		result = max(result, length)
	}

	return result
}
