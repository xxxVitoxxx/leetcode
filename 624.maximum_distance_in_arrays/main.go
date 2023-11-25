package main

import "fmt"

func main() {
	fmt.Println(maxDistance([][]int{{1, 4}, {0, 5}}))
}

func maxDistance(arrays [][]int) int {
	var res int
	minVal, maxVal := arrays[0][0], arrays[0][len(arrays[0])-1]
	for i := 1; i < len(arrays); i++ {
		res = max(res, abs(maxVal-arrays[i][0]))
		res = max(res, abs(arrays[i][len(arrays[i])-1]-minVal))
		maxVal = max(maxVal, arrays[i][len(arrays[i])-1])
		minVal = min(minVal, arrays[i][0])
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
