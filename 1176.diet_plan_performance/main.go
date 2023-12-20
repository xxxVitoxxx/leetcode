package main

// brute force
func dietPlanPerformance(calories []int, k, lower, upper int) int {
	var ans int
	for pnt1 := 0; pnt1 < len(calories); pnt1++ {
		pnt2 := pnt1 + k
		if pnt2 <= len(calories) {
			var consume int
			for i := pnt1; i < pnt2; i++ {
				consume += calories[i]
			}

			if consume < lower {
				ans--
			}
			if consume > upper {
				ans++
			}
		}
	}
	return ans
}

func dietPlanPerformance2(calories []int, k, lower, upper int) int {
	prefixSum, prev := make([]int, len(calories)), 0
	for i := 0; i < len(calories); i++ {
		prefixSum[i] = calories[i] + prev
		prev = prefixSum[i]
	}

	ans, prev := 0, 0
	for i := k; i <= len(prefixSum); i++ {
		if prefixSum[i-1]-prev < lower {
			ans--
		}
		if prefixSum[i-1]-prev > upper {
			ans++
		}
		prev = prefixSum[i-k]
	}
	return ans
}

func dietPlanPerformance3(calories []int, k, lower, upper int) int {
	ans, curr := 0, 0
	for i := 0; i < len(calories); i++ {
		curr += calories[i]
		if i-k >= 0 {
			curr -= calories[i-k]
		}
		if i+1-k >= 0 {
			if curr < lower {
				ans--
			}
			if curr > upper {
				ans++
			}
		}
	}
	return ans
}
