package main

func getDistances(arr []int) []int64 {
	idxMap := make(map[int][]int)
	idxSumMap := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		idxMap[arr[i]] = append(idxMap[arr[i]], i)
		idxSumMap[arr[i]] += i
	}

	ans := make([]int64, len(arr))
	for k, indexes := range idxMap {
		var leftSum int
		for i, idx := range indexes {
			ans[idx] = int64(idx*i-leftSum) + int64(idxSumMap[k]-leftSum-idx-idx*(len(indexes)-i-1))
			leftSum += idx
		}
	}
	return ans
}

func getDistances2(arr []int) []int64 {
	numsMap := make(map[int][]int)
	sumMap := make(map[int]int)
	for i, num := range arr {
		numsMap[num] = append(numsMap[num], i)
		sumMap[num] += i
	}

	ans := make([]int64, len(arr))
	for k, nums := range numsMap {
		var prev int
		for i, num := range nums {
			delta := num - prev
			prev = num
			sumMap[k] += delta * (2*i - len(nums))
			ans[num] = int64(sumMap[k])
		}
	}
	return ans
}

func getDistances3(arr []int) []int64 {
	m := make(map[int][]int)
	for i := 0; i < len(arr); i++ {
		m[arr[i]] = append(m[arr[i]], i)
	}

	ans := make([]int64, len(arr))
	for _, v := range m {
		var diff int
		for i := 0; i < len(v); i++ {
			diff += v[i] - v[0]
		}

		ans[v[0]] = int64(diff)
		for i := 1; i < len(v); i++ {
			delta := v[i] - v[i-1]
			diff += delta * (2*i - len(v))
			ans[v[i]] = int64(diff)
		}
	}
	return ans
}
