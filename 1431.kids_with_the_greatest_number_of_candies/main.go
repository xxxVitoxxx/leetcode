package main

// time complexity: O(n)
// space complexity: O(n)
func kidsWithCandies(candies []int, extraCandies int) []bool {
	// slices.Max(candies)
	var maxCandy int
	for _, candy := range candies {
		if candy > maxCandy {
			maxCandy = candy
		}
	}

	ans := make([]bool, len(candies))
	for i := 0; i < len(candies); i++ {
		ans[i] = candies[i]+extraCandies >= maxCandy
	}
	return ans
}
