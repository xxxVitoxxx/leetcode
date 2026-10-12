package main

func countComponents(n int, edges [][]int) int {
	adjList := make(map[int][]int)
	for _, edge := range edges {
		adjList[edge[0]] = append(adjList[edge[0]], edge[1])
		adjList[edge[1]] = append(adjList[edge[1]], edge[0])
	}

	visited := make(map[int]bool)
	var hasCycle func(int, int)
	hasCycle = func(child, parent int) {
		visited[child] = true

		for _, nb := range adjList[child] {
			if nb == parent || visited[nb] {
				continue
			}

			hasCycle(nb, child)
		}
	}

	var count int
	for i := range n {
		if visited[i] {
			continue
		}

		count++

		hasCycle(i, -1)
	}

	return count
}

// Disjoint Set Union
func countComponents2(n int, edges [][]int) int {
	uf := newUnionFind(n)
	for _, edge := range edges {
		uf.union(edge[0], edge[1])
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
		rank[i] = 0
	}
	return &unionFind{root, rank, n}
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
