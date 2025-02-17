package main

// recursion
// time complexity: O(2^n). this is the slowest way to solve the
// fibonacci sequence because it takes exponential time. the amount
// of operations needed, for each level of recursion, grows exponentially
// as the depth approaches n.
//
// space complexity: O(n). we need space proportional to n to account for
// the max size of the stack, in memory. this stack keeps track of the
// function
//
// use recursion to compute the fibonacci number of a given number
//   - check if the provider input value, n is less than or equal 1
//     if true, return n.
//   - otherwise, the function fib() calls itself, with the result
//     of the 2 previous number being added to each other, passed
//     in as the argument.
//     this is derived directly from the recurrence relation:
//     fib(n) = fib(n-1) + fib(n-2)
//   - do this until all numbers have been computed, then return the
//     resulting answer.
func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

var memo = make(map[int]int)

// top-down approach using memoization
// time complexity: O(n). each number is visited, computed and
// then stored for O(1) access later on.
//
// space complexity: O(n). the size of the stack in memory is
// proportional to n. also, the memoization hash table is used,
// which occupies O(n) space.
func fib2(n int) int {
	if n < 2 {
		return n
	}
	if val, ok := memo[n]; ok {
		return val
	}
	memo[n] = fib(n-1) + fib(n-2)
	return memo[n]
}

// bottom-up approach using tabulation
// time complexity: O(n)
// each number is visited, computed and then stored for O(1)
// access later on.
//
// space complexity: O(n)
// the size of the data structure is proportional to n.
//
// improve upon the recursive approach by using iteration, still
// solving for all of the sub-problems and returning the answer
// for n, using already computed Fibonacci values. while using
// a bottom-up approach, we can iteratively compute and store
// the values, only returning once we reach the result.
//   - if n is less than or equal 1, return n
//   - otherwise, iterate through n, storing each computed answer
//     in an array along the way.
//   - use this array as a reference to the 2 previous numbers to
//     calculate the current Fibonacci number.
//   - once we've reached the last number, return its Fibonacci number.
func fib3(n int) int {
	if n < 2 {
		return n
	}
	memo := make([]int, n+1)
	memo[1] = 1
	for i := 2; i <= n; i++ {
		memo[i] = memo[i-1] + memo[i-2]
	}
	return memo[n]
}

// iterative bottom-up approach
// time complexity: O(n)
// each value from 2 to n is computed once. thus, the time it takes to
// find the answer is directly proportional to n where n is the Fibonacci
// Number we are looking to compute.
//
// space complexity: O(1)
// this requires 1 unit of space for the integer n and 2 units of space to
// store the computed values(first and second) for every loop iteration.
// the amount of space used in independent of n, so this approach use a
// constant amount of space.
//
// let's get rid of the need to use all of that space and instead
// use the minimum amount of space required. notice that during
// each recursive call in the top-down approach and each iteration
// in the bottom-up approach, we only needed to look at the results
// of fib(n-1) and fib(n-2) to determine the result of fib(n).
// therefore, we can achieve O(1) space complexity by only storing
// the value of the two previous numbers and updating them as we
// iterate to n.
//   - if n is less than or equal 1, return n
//   - we need 2 variables to store each state fib(n-1) and fib(n-2)
//   - preset the initial values:
//     ** initialize first with 0, since this will represent fib(n-2)
//     when computing the current value
//     ** initialize second with 1, since this will represent fib(n-1)
//     when computing the current value
//   - iterate incrementally by 1, all the way up to and including n.
//     starting at 2, since 0 and 1 are pre-computed
//   - set the first value to second
//   - set the second value to first+second because that is the value
//     we are currently computing
//   - when we reach n+1, we will exit the loop and return the previously
//     set second value
func fib4(n int) int {
	if n < 2 {
		return n
	}
	first, second := 0, 1
	for i := 2; i <= n; i++ {
		first, second = second, first+second
	}
	return second
}
