package main

// brute force will encounter a time limit
// time complexity: O(2^n)
// space complexity: O(n)
//
// in this approach we take all possible step combinations
// i.e. 1 and 2, at every step. at every step we are calling
// the helper function for step 1 and step 2, and return the
// sum of returned values of both functions.
func climbStairs(n int) int {
	return helper(0, n)
}

func helper(i, n int) int {
	if i > n {
		return 0
	}
	if i == n {
		return 1
	}
	return helper(i+1, n) + helper(i+2, n)
}

// recursion with memoization
// time complexity: O(n). size of recursion tree can go up to n
// space complexity: O(n). the depth of recursion tree can go up to n
//
// in the previous approach we are redundantly calculating the result
// for every step. instead, we can store the result at each step in
// memo array and directly returning the result from the memo array
// whenever that function is called again.
//
// in this way we are pruning recursion tree with the help of memo array
// and reducing the size of recursion tree up to n.
func climbStairs2(n int) int {
	memo := make([]int, n+1)
	return helper2(0, n, memo)
}

func helper2(i, n int, memo []int) int {
	if i > n {
		return 0
	}
	if i == n {
		return 1
	}
	if memo[i] > 0 {
		return memo[i]
	}
	memo[i] = helper2(i+1, n, memo) + helper2(i+2, n, memo)
	return memo[i]
}

// dynamic programming
// time complexity: O(n). single loop up to n
// space complexity: O(n). dp array of size n is used
//
// as we can see this problem can be broken into sub problems,
// and it contains the optimal substructure property i.e. its
// optimal solution can be constructed efficiently from optimal
// solutions of its sub problems, we can use dynamic programming
// to solve this problem.
//
// one can reach i^th step in one of the two ways:
//  1. Taking a single step from (i-1)^th step.
//  2. Taking a step of 2 from (i-2)^th step.
//
// so, the total number of ways to reach i^th is equal to sum of
// ways of reaching (i-1)^th step and ways of reaching (i-2)^th step.
//
// let dp[i] denotes the number of ways to reach on i^th step:
// dp[i] = dp[i-1] + dp[i-2]
//
// dp [0, 1, 2, 3, 5, 8, 13 ...]
func climbStairs3(n int) int {
	if n <= 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// fibonacci number
// time complexity: O(n). single loop up to n is required
// to calculate n^th fibonacci number.
//
// space complexity: O(1). constant space is used.
//
// in the above approach we have used dp array where
// dp[i] = dp[i-1] + dp[i-2]. it can be easily analysed
// that dp[i] is nothing bu i^th fibonacci number.
//
// Fib(n) = Fib(n-1) + Fib(n-2)
//
// Now we just have to find n^th number of the fibonacci series
// having 1 and 2 their first and second term respectively,
// i.e. Fib(1) = 1 and Fib(2) = 2.
func climbStairs4(n int) int {
	if n <= 1 {
		return n
	}
	first, second := 1, 2
	for i := 3; i <= n; i++ {
		first, second = second, first+second
	}
	return second
}
