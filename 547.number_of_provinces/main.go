package main

// dfs
// time complexity: O(n^2)
// space complexity: O(n)
func findCircleNum(isConnected [][]int) int {
	visited := make(map[int]bool, len(isConnected))
	var dfs func(int)
	dfs = func(node int) {
		visited[node] = true
		for i := 0; i < len(isConnected); i++ {
			if isConnected[node][i] == 1 && !visited[i] {
				dfs(i)
			}
		}
	}

	var provinces int
	for i := 0; i < len(isConnected); i++ {
		if !visited[i] {
			provinces++
			dfs(i)
		}
	}

	return provinces
}

// bfs
// time complexity: O(n)
// space complexity: O(n)
func findCircleNum2(isConnected [][]int) int {
	visited := make(map[int]bool, len(isConnected))
	queue := []int{}
	provinces := 0
	for i := 0; i < len(isConnected); i++ {
		if !visited[i] {
			provinces++
			queue = append(queue, i)
		}

		for len(queue) > 0 {
			curr := queue[0]
			queue = queue[1:]
			visited[curr] = true
			for j := 0; j < len(isConnected); j++ {
				if isConnected[curr][j] == 1 && !visited[j] {
					queue = append(queue, j)
				}
			}
		}
	}

	return provinces
}

func findCircleNum3(isConnected [][]int) int {
	uf := newUnionFind(len(isConnected))
	for i, conn := range isConnected {
		for j := range conn {
			if conn[j] != 0 {
				uf.union(i, j)
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

func newUnionFind(n int) *unionFind {
	root, rank := make([]int, n), make([]int, n)
	for i := range n {
		root[i] = i
		rank[i] = 1
	}
	return &unionFind{root, rank, n}
}

func (uf *unionFind) find(x int) int {
	if x == uf.root[x] {
		return x
	}

	uf.root[x] = uf.find(uf.root[x])
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

func (uf *unionFind) getCount() int {
	return uf.count
}
