package main

// DFS
// use hash table to record already visited.
// time complexity: O(MN)
// space complexity: O(MN)
func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	visited := make(map[[2]int]bool)

	directions := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	var dfs func(int, int)
	dfs = func(row, col int) {
		visited[[2]int{row, col}] = true

		for _, dir := range directions {
			newRow, newCol := row+dir[0], col+dir[1]
			if newRow < 0 || newCol < 0 || newRow >= m || newCol >= n || visited[[2]int{newRow, newCol}] || grid[newRow][newCol] == '0' {
				continue
			}

			dfs(newRow, newCol)
		}
	}

	var result int
	for i := range m {
		for j := range n {
			if visited[[2]int{i, j}] || grid[i][j] == '0' {
				continue
			}

			result++
			dfs(i, j)
		}
	}

	return result
}

// DFS
// use list to record visited. this solution will faster than use hash table.
// time complexity: O(MN)
// space complexity: O(MN)
func numIslands2(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	directions := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	var dfs func(int, int)
	dfs = func(row, col int) {
		visited[row][col] = true

		for _, dir := range directions {
			newRow, newCol := row+dir[0], col+dir[1]
			if newRow < 0 || newCol < 0 || newRow >= m || newCol >= n || visited[newRow][newCol] || grid[newRow][newCol] == '0' {
				continue
			}

			dfs(newRow, newCol)
		}
	}

	var result int
	for i := range m {
		for j := range n {
			if visited[i][j] || grid[i][j] == '0' {
				continue
			}

			result++
			dfs(i, j)
		}
	}

	return result
}

// DFS
// time complexity: O(MN)
// space complexity: O(MN)
func numIslands3(grid [][]byte) int {
	m, n := len(grid), len(grid[0])

	var dfs func(int, int)
	dfs = func(row, col int) {
		if row < 0 || col < 0 || row >= m || col >= n || grid[row][col] != '1' {
			return
		}

		grid[row][col] = '0'

		dfs(row+1, col)
		dfs(row-1, col)
		dfs(row, col+1)
		dfs(row, col-1)
	}

	var result int
	for i := range m {
		for j := range n {
			if grid[i][j] != '1' {
				continue
			}

			result++
			dfs(i, j)
		}
	}

	return result
}

// BFS
// time complexity: O(MN)
// space complexity: O(min(M, N))
func numIslands4(grid [][]byte) int {
	directions := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	var result int
	m, n := len(grid), len(grid[0])
	for i := range m {
		for j := range n {
			if grid[i][j] != '1' {
				continue
			}

			result++

			neighbors := [][]int{{i, j}}
			grid[i][j] = '0'
			for len(neighbors) > 0 {
				curr := neighbors[0]
				neighbors = neighbors[1:]

				for _, dir := range directions {
					newRow, newCol := curr[0]+dir[0], curr[1]+dir[1]
					if newRow < 0 || newCol < 0 || newRow >= m || newCol >= n || grid[newRow][newCol] != '1' {
						continue
					}

					neighbors = append(neighbors, []int{newRow, newCol})
					grid[newRow][newCol] = '0'
				}
			}
		}
	}

	return result
}

// disjoint set
// time complexity: O(MN)
// space complexity: O(MN)
func numIslands5(grid [][]byte) int {
	uf := newUnionFind(grid)

	m, n := len(grid), len(grid[0])
	for i := range m {
		for j := range n {
			if grid[i][j] != '1' {
				continue
			}

			if i+1 < m && grid[i+1][j] == '1' {
				uf.union(i*n+j, (i+1)*n+j)
			}

			if j+1 < n && grid[i][j+1] == '1' {
				uf.union(i*n+j, i*n+j+1)
			}
		}
	}

	return uf.count
}

type unionFind struct {
	root  []int
	rank  []int
	count int
}

func newUnionFind(grid [][]byte) *unionFind {
	m, n := len(grid), len(grid[0])
	root, rank := make([]int, m*n), make([]int, m*n)
	var count int
	for i := range m {
		for j := range n {
			if grid[i][j] != '1' {
				continue
			}

			count++
			root[i*n+j] = i*n + j
		}
	}
	return &unionFind{root, rank, count}
}

func (uf *unionFind) find(x int) int {
	if x != uf.root[x] {
		uf.root[x] = uf.find(uf.root[x])
	}
	return uf.root[x]
}

func (uf *unionFind) union(x, y int) {
	rootX, rootY := uf.find(x), uf.find(y)
	if rootX != rootY {
		if uf.rank[rootX] > uf.rank[rootY] {
			uf.root[rootY] = rootX
		} else if uf.rank[rootX] < uf.rank[rootY] {
			uf.root[rootX] = rootY
		} else {
			uf.root[rootY] = rootX
			uf.rank[rootX]++
		}

		uf.count--
	}
}
