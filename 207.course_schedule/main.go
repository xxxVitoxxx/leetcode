package main

// Topological Sort Using Kahn's Algorithm(BFS)
//
// n is the number of courses and m is the size of prerequisites
// time complexity: O(n+m)
// space complexity: O(n+m)
func canFinish(numCourses int, prerequisites [][]int) bool {
	adjList := make(map[int][]int)
	inDegree := make([]int, numCourses)
	for _, pre := range prerequisites {
		adjList[pre[1]] = append(adjList[pre[1]], pre[0])
		inDegree[pre[0]]++
	}

	var queue []int
	for i, d := range inDegree {
		if d == 0 {
			queue = append(queue, i)
		}
	}

	var count int
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		count++

		for _, course := range adjList[curr] {
			inDegree[course]--
			if inDegree[course] == 0 {
				queue = append(queue, course)
			}
		}
	}

	return count == numCourses
}

// Topological Sort Using Kahn's Algorithm(DFS)
//
// n is the number of courses and m is the size of prerequisites
// time complexity: O(n+m)
// space complexity: O(n+m)
func canFinish2(numCourses int, prerequisites [][]int) bool {
	adjList := make(map[int][]int)

	for _, pre := range prerequisites {
		adjList[pre[1]] = append(adjList[pre[1]], pre[0])
	}

	// value description
	// 0 -> not visited yet
	// 1 -> recursive call in progress
	// 2 -> recursive call finished and no cycle found
	states := make([]int, numCourses)

	var hasCycle func(n int) bool
	hasCycle = func(n int) bool {
		// 走回曾經訪問過的節點，表示有cycle
		if states[n] == 1 {
			return true
		}

		// 該節點已訪問過且確認沒有 cycle
		if states[n] == 2 {
			return false
		}

		states[n] = 1
		for _, course := range adjList[n] {
			if hasCycle(course) {
				return true
			}
		}

		states[n] = 2
		return false
	}

	for i := range numCourses {
		if hasCycle(i) {
			return false
		}
	}

	return true
}
