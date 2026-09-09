package main

// time complexity: O(logN)
// if num was even, we halved what was left. If it was odd, we only subtracted 1.
// however, by subtracting 1, we were making it even, and so on the next step we were guaranteed to halve it.
// what this means is that in the worst case, we're halving it on every second step.
// we treat the 1/2 of the time as a constant though, so in essence, we say that at each step, num is being halved.
// when something is halved at every step, it has a O(logN) time complexity.
//
// space complexity: O(1)
// we only use a constant number of integer variables, and so the space complexity is O(1).
func numberOfSteps1(num int) int {
	var step int
	for num > 0 {
		// even
		if num%2 == 0 {
			num /= 2
		} else {
			// odd
			num--
		}
		step++
	}
	return step
}

// counting bits
// do not change the time complexity, but they offer a different way of thinking about the problem
//
// `>> and & bit operation`
// 4(100) >> 1 = 2(10)
// 5(101) &  1 = 1
//
// - if the last bit is 1, we need two operations: subtract 1 and divide by 2.
// - if the last bit is 0, we need one operations: divide by 2.
//
// we repeat this process while the number has a least 2 bits.
// the final case when the number becomes 1 is handled separately by
// adding one last step at the return statement.

// time complexity: O(logN)
// space complexity: O(1)
func numberOfSteps2(num int) int {
	var step int
	for num > 1 {
		step += 1 + num&1
		num >>= 1
	}
	return step + num
}
