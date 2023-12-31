package main

// prefix sum
// 拆成兩個問題，假設 index i 的數字為 num ，則 index 的答案為以下兩項的加總
// 1. The sum of absolute differences between num and all numbers less than num.
//    leftTotal = leftCount * nums[i] - leftSum
// 2. The sum of absolute differences between num and all numbers greater than num.
//    rightTotal = rightSum - (len(nums)-i-1) * nums[i]

func getSumAbsoluteDifferences(nums []int) []int {
	prefix := make([]int, len(nums))
	prefix[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		prefix[i] = prefix[i-1] + nums[i]
	}

	// calculate the sum of absolute differences for each element
	ans := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		leftSum := prefix[i] - nums[i]
		rightSum := prefix[len(prefix)-1] - prefix[i]

		leftTotal := i*nums[i] - leftSum
		rightTotal := rightSum - (len(nums)-i-1)*nums[i]
		ans[i] = leftTotal + rightTotal
	}
	return ans
}

func getSumAbsoluteDifferences2(nums []int) []int {
	var sum int
	for _, num := range nums {
		sum += num
	}

	var leftSum int
	ans := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		rightSum := sum - leftSum - nums[i]

		leftTotal := i*nums[i] - leftSum
		rightTotal := rightSum - (len(nums)-i-1)*nums[i]

		ans[i] = leftTotal + rightTotal
		leftSum += nums[i]
	}
	return ans
}

func getSumAbsoluteDifferences3(nums []int) []int {
	var diffSum int
	for _, num := range nums {
		diffSum += num - nums[0]
	}

	ans := make([]int, len(nums))
	ans[0] = diffSum
	for i := 1; i < len(nums); i++ {
		delta := nums[i] - nums[i-1]
		diffSum += delta*i - delta*(len(nums)-i)
		ans[i] = diffSum
	}
	return ans
}

func getSumAbsoluteDifferences4(nums []int) []int {
	var sum int
	for _, num := range nums {
		sum += num
	}

	ans := make([]int, len(nums))
	var prev int
	for i := 0; i < len(nums); i++ {
		delta := nums[i] - prev
		prev = nums[i]
		sum += delta * (2*i - (len(nums)))
		ans[i] = sum
	}
	return ans
}
