package main

import (
	"strconv"
	"strings"
)

// brute force
// time complexity: O(n * n * n)
// space complexity: O(1)
func equalPairs(grid [][]int) int {
	var answer int
	n := len(grid)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			var k int
			for ; k < n; k++ {
				if grid[i][k] != grid[k][j] {
					break
				}
			}
			if k == n {
				answer++
			}
		}
	}
	return answer
}

// hash
// time complexity: O(n * n)
// space complexity: O(n)
func equalPairs2(grid [][]int) int {
	n := len(grid)
	rowStrings := make(map[string]int)

	// build string for rows
	for _, row := range grid {
		var rowBuilder strings.Builder
		for _, num := range row {
			rowBuilder.WriteString(strconv.Itoa(num))
			rowBuilder.WriteString(",")
		}

		rowStr := rowBuilder.String()
		rowStrings[rowStr]++
	}

	var count int
	for col := 0; col < n; col++ {
		var colBuilder strings.Builder
		for row := 0; row < n; row++ {
			colBuilder.WriteString(strconv.Itoa(grid[row][col]))
			colBuilder.WriteString(",")
		}

		colStr := colBuilder.String()
		count += rowStrings[colStr]
	}

	return count
}

// trie
// time complexity: O(n * n)
// space complexity: O(n * n) in the worst case the trie has n * n nodes.
func equalPairs3(grid [][]int) int {
	root := newTrie()
	for _, row := range grid {
		root.insert(row)
	}

	n := len(grid)
	count := 0
	for col := 0; col < n; col++ {
		nums := make([]int, n)
		for row := 0; row < n; row++ {
			nums[row] = grid[row][col]
		}

		count += root.search(nums)
	}

	return count
}

type Trie struct {
	count    int
	children map[int]*Trie
}

func newTrie() *Trie {
	return &Trie{children: make(map[int]*Trie)}
}

func (t *Trie) insert(nums []int) {
	node := t
	for _, num := range nums {
		if _, ok := node.children[num]; !ok {
			node.children[num] = newTrie()
		}
		node = node.children[num]
	}

	node.count++
}

func (t *Trie) search(nums []int) int {
	node := t
	for _, num := range nums {
		if _, ok := node.children[num]; !ok {
			return 0
		}
		node = node.children[num]
	}

	return node.count
}
