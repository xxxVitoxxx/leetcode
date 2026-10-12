package main

import (
	"sort"
	"strings"
)

// Disjoint Set Union
func smallestStringWithSwaps(s string, pairs [][]int) string {
	uf := newUnionFind(len(s))
	for _, pair := range pairs {
		uf.union(pair[0], pair[1])
	}

	// 先找出所有 root 群有哪些文字
	m := make(map[int][]byte)
	for i := range s {
		root := uf.find(i)
		m[root] = append(m[root], s[i])
	}

	// 由小到大排序 root 群的文字
	for _, v := range m {
		sort.Slice(v, func(i, j int) bool {
			return v[i] < v[j]
		})
	}

	// 依序拿 root 最小的文字
	var str strings.Builder
	for i := range s {
		root := uf.find(i)
		str.WriteByte(m[root][0])
		m[root] = m[root][1:]
	}

	return str.String()
}

type unionFind struct {
	root []int
	rank []int
}

func newUnionFind(n int) *unionFind {
	root, rank := make([]int, n), make([]int, n)
	for i := range n {
		root[i] = i
		rank[i] = 0
	}
	return &unionFind{root, rank}
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
	}
}
