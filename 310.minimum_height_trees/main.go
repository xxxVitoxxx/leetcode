package main

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	adjMatrix := make([][]int, n)
	degree := make([]int, n)
	for _, edge := range edges {
		adjMatrix[edge[0]] = append(adjMatrix[edge[0]], edge[1])
		adjMatrix[edge[1]] = append(adjMatrix[edge[1]], edge[0])
		degree[edge[0]]++
		degree[edge[1]]++
	}

	var leaf []int
	for i, d := range degree {
		if d == 1 {
			leaf = append(leaf, i)
		}
	}

	remaining := n
	// 剩下兩個節點，就只能是兩個互相連在一起的兩個點
	// 這時兩個節點的度數都是 1 ，兩個都是葉子
	for remaining > 2 {
		remaining -= len(leaf)

		var newLeaf []int
		for i := range leaf {
			for _, nb := range adjMatrix[leaf[i]] {
				degree[nb]--
				if degree[nb] == 1 {
					newLeaf = append(newLeaf, nb)
				}
			}
		}

		leaf = newLeaf
	}

	return leaf
}
