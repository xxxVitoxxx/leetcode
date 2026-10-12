package main

// DFS
// V is the number of vertex, and E is the number of edges.
// time complexity: O(V+E)
// space complexity: O(V+E)
func validTree(n int, edges [][]int) bool {
	if len(edges) != n-1 {
		return false
	}

	adjList := make(map[int][]int)
	for _, edge := range edges {
		adjList[edge[0]] = append(adjList[edge[0]], edge[1])
		adjList[edge[1]] = append(adjList[edge[1]], edge[0])
	}

	visited := make(map[int]bool)
	var hasCycle func(int, int) bool
	hasCycle = func(child, parent int) bool {
		visited[child] = true

		for _, nb := range adjList[child] {
			if nb == parent {
				continue
			}

			if visited[nb] {
				return true
			}

			if hasCycle(nb, child) {
				return true
			}
		}

		return false
	}

	for i := range n {
		if visited[i] {
			continue
		}

		if hasCycle(i, -1) {
			return false
		}
	}

	return true
}

// Disjoint Set Union
// V is the number of vertex
// time complexity: O(V⋅α(V))
// space complexity: O(V)
func validTree2(n int, edges [][]int) bool {
	if len(edges) != n-1 {
		return false
	}

	uf := newUnionFind(n)
	for _, edge := range edges {
		if !uf.union(edge[0], edge[1]) {
			return false
		}
	}
	return true
}

type unionFind struct {
	root []int
	rank []int
}

func newUnionFind(n int) *unionFind {
	root, rank := make([]int, n), make([]int, n)
	for i := range n {
		root[i] = i
		rank[i] = 1
	}
	return &unionFind{root, rank}
}

func (uf *unionFind) find(x int) int {
	if x != uf.root[x] {
		uf.root[x] = uf.find(uf.root[x])
	}
	return uf.root[x]
}

func (uf *unionFind) union(x, y int) bool {
	rootX, rootY := uf.find(x), uf.find(y)
	if rootX == rootY {
		return false
	}

	if uf.rank[rootX] > uf.rank[rootY] {
		uf.root[rootY] = rootX
	} else if uf.rank[rootX] < uf.rank[rootY] {
		uf.root[rootX] = rootY
	} else {
		uf.root[rootY] = rootX
		uf.rank[rootX]++
	}

	return true
}
