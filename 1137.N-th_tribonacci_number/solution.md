# Solution

## Dynamic Programming(Top Down)

Let `dfs(i)` be the value of the `iᵗʰ` tribonacci number, and according to the given recurrence relation, we have:  
`dfs(i) = dfs(i-1) + dfs(i-2) + df(i-3 )`.  

Note that the solution to the current problem `dfs(i)` can be built from the solutions to its subproblem (`dfs(i-1)`, `dfs(i-2)`, `dfs(i-3)`). Dynamic Programming is exactly based on the concept of **overlapping subproblem** and **optimal substructure**, thus it can be a powerful tool to solve this problem efficiently.  

> For example, suppose we are given `n=5`. we have `dfs(5) = dfs(4) + dfs(3) + dfs(2)` according to the recursion relation.  
>
> However, we don't know the values of `dfs(4)` or `dfs(3)`, so we need to continue referring to the recursion relation, first on `dfs(4)`: `dfs(4) = dfs(3) + dfs(2) + dfs(1)`.  
>
> We don't know the value of `dfs(3)` either, continue referring to the recursion relation: `dfs(3) = dfs(2) + dfs(1) + dfs(0)`.  
>
> Great, now we have some known solutions, as the values of the first three tribonacci numbers are given in the problem description:  
> - `dfs(0) = 0`
> - `dfs(1) = 1`
> - `dfs(2) = 1`
>
> With these base cases, we can find `dfs(3)`. Once we have `dfs(3) = 2`, this allows us to find the answer of `dfs(4)`, and then finally `dfs(5)`.  

In general, we recursively break the current problem down into subproblem, until we reach the base cases: the first three tribonacci numbers.  

For these base cases, we don't need further recursion. for any other `i>2`, we can refer to the recurrence relation above. since the subproblem always had a smaller `i` than the current problem, we are guaranteed to eventually reach base cases.  

However, we notice that the same `dfs(i)` may be calculated multiple times. to avoid the high time complexity caused by repeated calculations, we can use a hash map `dp` to save results. this is a technique called memoization.  

![](dp1.png)  

In later calculations, if we find that `dp[i]` already exists, we know that we have already computed the value of `dfs(i)`, and we can simply return the precomputed solution `dp[i]` without further recursion. this can significantly reduce the computational time and make the algorithm more efficient.  
