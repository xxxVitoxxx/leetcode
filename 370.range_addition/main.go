package main

// Diff Array
// time complexity: O(M+N)
// space complexity: O(1)
func getModifiedArray(length int, updates [][]int) []int {
	arr := make([]int, length+1)
	for _, update := range updates {
		arr[update[0]] += update[2]
		arr[update[1]+1] -= update[2]
	}

	for i := 1; i < length; i++ {
		arr[i] += arr[i-1]
	}

	return arr[:length]
}
