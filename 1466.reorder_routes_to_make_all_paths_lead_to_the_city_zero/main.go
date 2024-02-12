package main

// dfs
// time complexity: O(n)
// space complexity: O(n)
func minReorder(n int, connections [][]int) int {
	adList := make(map[int][]int, n)
	for _, conn := range connections {
		// append the outgoing road to the city.
		adList[conn[0]] = append(adList[conn[0]], conn[1])
		// append incoming road to the city.
		// we need a way to differentiate between incoming and outgoing road,
		// so we can mark this edge negative.
		adList[conn[1]] = append(adList[conn[1]], -conn[0])
	}

	visited := make(map[int]bool, n)
	var count int
	var dfs func(int)
	dfs = func(num int) {
		visited[num] = true
		for _, nb := range adList[num] {
			if !visited[abs(nb)] {
				if nb > 0 {
					// outgoing edge that need to be reoriented
					count++
				}
				dfs(abs(nb))
			}
		}
	}
	dfs(0)

	return count
}

// bfs
// time complexity: O(n)
// space complexity: O(n)
func minReorder2(n int, connections [][]int) int {
	adList := make(map[int][]int, n)
	for _, conn := range connections {
		adList[conn[0]] = append(adList[conn[0]], conn[1])
		adList[conn[1]] = append(adList[conn[1]], -conn[0])
	}

	queue := []int{0}
	visited := make(map[int]bool, n)
	count := 0
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		visited[curr] = true
		for _, nb := range adList[curr] {
			if !visited[abs(nb)] {
				if nb > 0 {
					count++
					queue = append(queue, nb)
				} else {
					queue = append(queue, abs(nb))
				}
			}
		}
	}
	return count
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// dfs
// time complexity: O(n)
// space complexity: O(n)
func minReorder3(n int, connections [][]int) int {
	edges := make(map[int][][]int, n)
	for _, conn := range connections {
		edges[conn[0]] = append(edges[conn[0]], conn)
		edges[conn[1]] = append(edges[conn[1]], conn)
	}

	visited := make(map[int]bool, n)
	var count int
	var dfs func(int)
	dfs = func(num int) {
		visited[num] = true
		for _, edge := range edges[num] {
			src := edge[0]
			dest := edge[1]
			if !visited[dest] {
				count++
				dfs(dest)
			}
			if dest == num && !visited[src] {
				dfs(src)
			}
		}
	}
	dfs(0)

	return count
}

// bfs
// time complexity: O(n)
// space complexity: O(n)
func minReorder4(n int, connections [][]int) int {
	edges := make(map[int][][]int, n)
	for _, conn := range connections {
		edges[conn[0]] = append(edges[conn[0]], conn)
		edges[conn[1]] = append(edges[conn[1]], conn)
	}

	queue := []int{0}
	visited := make(map[int]bool, n)
	count := 0
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		visited[curr] = true
		for _, edge := range edges[curr] {
			src := edge[0]
			dest := edge[1]
			if !visited[dest] {
				count++
				queue = append(queue, dest)
			} else if dest == curr && !visited[src] {
				queue = append(queue, src)
			}
		}
	}

	return count
}
