package main

// stack
// time complexity: O(n)
// space complexity: O(n)
func asteroidCollision(asteroids []int) []int {
	var ans []int
	for _, asteroid := range asteroids {
		flag := true
		for len(ans) > 0 && ans[len(ans)-1] > 0 && asteroid < 0 {
			if ans[len(ans)-1] < -asteroid {
				ans = ans[:len(ans)-1]
				continue
			} else if ans[len(ans)-1] == -asteroid {
				ans = ans[:len(ans)-1]
			}
			// the current asteroid will be destroyed
			flag = false
			break
		}

		if flag {
			ans = append(ans, asteroid)
		}
	}
	return ans
}

// stack
// time complexity: O(n)
// space complexity: O(n)
func asteroidCollision2(asteroids []int) []int {
	var ans []int
	for i := 0; i < len(asteroids); i++ {
		if len(ans) == 0 || ans[len(ans)-1] < 0 || asteroids[i] > 0 {
			ans = append(ans, asteroids[i])
		} else {
			diff := ans[len(ans)-1] + asteroids[i]
			if diff == 0 {
				ans = ans[:len(ans)-1]
			} else if diff < 0 {
				// if diff is less than 0, continue to use the origin asteroids[i] to compare with
				// the last one of the ans.
				ans = ans[:len(ans)-1]
				i--
			}
		}
	}
	return ans
}
