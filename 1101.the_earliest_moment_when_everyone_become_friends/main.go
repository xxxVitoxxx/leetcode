package main

import "sort"

// Disjoint Set Union
func earliestAcq(logs [][]int, n int) int {
	sort.Slice(logs, func(i, j int) bool {
		return logs[i][0] < logs[j][0]
	})

	uf := newUnionFind(n)
	for _, log := range logs {
		if uf.union(log[1], log[2]) {
			n--
		}

		if n == 1 {
			return log[0]
		}
	}

	return -1
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
