package main

func deleteAndEarn(nums []int) int {
	var maxValue int
	pointMap := make(map[int]int)
	for _, num := range nums {
		maxValue = max(maxValue, num)
		pointMap[num] += num
	}

	dp := make([]int, maxValue+1)
	dp[1] = pointMap[1]

	for i := 2; i < len(dp); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+pointMap[i])
	}

	return dp[maxValue]
}

func deleteAndEarn2(nums []int) int {
	var maxValue int
	pointMap := make(map[int]int)
	for _, num := range nums {
		maxValue = max(maxValue, num)
		pointMap[num] += num
	}

	mem := make(map[int]int)

	var getMaxPoint func(int) int
	getMaxPoint = func(n int) int {
		if n == 0 {
			return 0
		}
		if n == 1 {
			return pointMap[1]
		}

		if value, ok := mem[n]; ok {
			return value
		}

		mem[n] = max(getMaxPoint(n-1), getMaxPoint(n-2)+pointMap[n])
		return mem[n]
	}

	return getMaxPoint(maxValue)
}

func deleteAndEarn3(nums []int) int {
	var maxValue int
	pointMap := make(map[int]int)
	for _, num := range nums {
		maxValue = max(maxValue, num)
		pointMap[num] += num
	}

	prev2, prev1 := 0, pointMap[1]
	for i := 2; i <= maxValue; i++ {
		prev2, prev1 = prev1, max(prev1, prev2+pointMap[i])
	}

	return prev1
}
